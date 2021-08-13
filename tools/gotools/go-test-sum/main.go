// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/erda-project/erda/tools/gotools/go-test-sum/cover"
)

var (
	codeFileExts = map[string]bool{
		".go":  true,
		".s":   true,
		".c":   true,
		".h":   true,
		".cpp": true,
	}
	// homeDir string
)

const testSumFilename = "go.test.sum"

// func init() {
// 	home, err := homedir.Dir()
// 	if err != nil {
// 		fmt.Println("failed to get home directory")
// 		os.Exit(1)
// 	}
// 	homeDir = home
// }

func main() {
	wd, err := os.Getwd()
	fmt.Println("--", wd, err)
	base, err := readBasePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = testAllPackages(base)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func testAllPackages(base string) error {
	packages := make(map[string]*packageInfo)
	fset := token.NewFileSet()
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip directories like ".git".
			if name := info.Name(); name != "." && strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			// parse package
			pkgs, err := parser.ParseDir(fset, path, nil, parser.ImportsOnly)
			if err != nil {
				return err
			}
			pkgPath := filepath.Join(base, path)
			pkgi := packages[pkgPath]
			if pkgi == nil {
				pkgi = &packageInfo{
					path:         path,
					dependencies: make(map[string]struct{}),
				}
				packages[pkgPath] = pkgi
			}
			for _, pkg := range pkgs {
				for _, info := range pkg.Files {
					for _, imp := range info.Imports {
						impPath, _ := strconv.Unquote(imp.Path.Value)
						pkgi.dependencies[impPath] = struct{}{}
					}
				}
			}
			return nil
		}
		pkgPath := filepath.Join(base, filepath.Dir(path))
		pkgi := packages[pkgPath]
		if pkgi == nil {
			return nil
		}
		// Only check code files.
		ext := filepath.Ext(path)
		if !codeFileExts[ext] {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			pkgi.hasTestFile = true
		}
		pkgi.files = append(pkgi.files, &fileInfo{
			path:       path,
			updateTime: info.ModTime(),
		})
		return nil
	})
	if err != nil {
		return err
	}

	// hash sum
	now := time.Now()
	pkgSum := make(map[string]*testSumItem)
	incoming := make(map[string]map[string]struct{})
	for pkg, info := range packages {
		if len(info.files) <= 0 {
			continue
		}
		sort.Sort(fileInfos(info.files))
		buf := make([]byte, 8)
		hash := md5.New()
		for _, file := range info.files {
			byts, err := ioutil.ReadFile(file.path)
			if err != nil {
				return err
			}
			hash.Write(byts)
			hash.Write(buf)
		}
		byts := hash.Sum(nil)
		pkgSum[pkg] = &testSumItem{
			pkg:      pkg,
			hash:     hex.EncodeToString(byts),
			testTime: now,
			info:     info,
		}

		for d := range info.dependencies {
			m := incoming[d]
			if m == nil {
				m = make(map[string]struct{})
				incoming[d] = m
			}
			m[pkg] = struct{}{}
		}
	}

	const doTestCheck = true
	if doTestCheck {
		preSum := readTestSum()
		profiles, err := cover.ParseProfiles("coverage.txt")
		if preSum == nil || err != nil {
			_, err := runTest("./...")
			if err != nil {
				return err
			}
		} else {
			coverage := make(map[string][]*cover.Profile)
			for pkg, sum := range pkgSum {
				if sum.tested {
					continue
				}
				pre, ok := preSum[pkg]
				if ok && pre.hash == sum.hash {
					sum.testTime = pre.testTime
					continue
				}
				err := recursiveTest(sum, pkgSum, incoming, coverage)
				if err != nil {
					return err
				}
			}
			var newProfiles []*cover.Profile
			for _, p := range profiles {
				dir := filepath.Dir(p.FileName)
				if _, ok := coverage[dir]; !ok {
					newProfiles = append(newProfiles, p)
				}
			}
			for _, ps := range coverage {
				newProfiles = append(newProfiles, ps...)
			}
			err = cover.Write("coverage.txt.tmp", "atomic", newProfiles)
			if err != nil {
				return err
			}
			os.Rename("coverage.txt.tmp", "coverage.txt")
		}
	}
	return writeTestSum(pkgSum)
}

type fileInfo struct {
	path       string
	updateTime time.Time
}

type fileInfos []*fileInfo

func (fs fileInfos) Len() int           { return len(fs) }
func (fs fileInfos) Less(i, j int) bool { return fs[i].path < fs[j].path }
func (fs fileInfos) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

type packageInfo struct {
	path         string
	files        []*fileInfo
	dependencies map[string]struct{}
	hasTestFile  bool
}

type testSumItem struct {
	pkg      string
	hash     string
	testTime time.Time
	info     *packageInfo
	tested   bool
}

func readTestSum() map[string]*testSumItem {
	byts, err := ioutil.ReadFile(testSumFilename)
	if err != nil {
		return nil
	}
	sum := make(map[string]*testSumItem)
	lines := strings.Split(string(byts), "\n")
	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			if len(fields) == 0 {
				continue
			}
			return nil
		}
		t, err := time.Parse(time.RFC3339Nano, fields[2])
		if err != nil {
			return nil
		}
		pkg := fields[0]
		sum[pkg] = &testSumItem{
			pkg:      pkg,
			hash:     fields[1],
			testTime: t,
		}
	}
	return sum
}

func writeTestSum(testSum map[string]*testSumItem) error {
	var keys []string
	for pkg := range testSum {
		keys = append(keys, pkg)
	}
	sort.Strings(keys)
	buf := bytes.Buffer{}
	for _, key := range keys {
		sum := testSum[key]
		buf.WriteString(fmt.Sprintf("%s %s %s\n", sum.pkg, sum.hash, sum.testTime.Format(time.RFC3339Nano)))
	}
	return ioutil.WriteFile(testSumFilename, buf.Bytes(), os.ModePerm)
}

func readBasePath() (string, error) {
	mod, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(mod), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(line[len("module "):]), nil
		}
	}
	return "", errors.New("not found module path")
}

func runTest(file string) (profiles []*cover.Profile, err error) {
	var coverage = "coverage.txt"
	if file != "./..." {
		coverage = filepath.Join(file, "pkg.coverage.txt")
		defer func() {
			profiles, _ = cover.ParseProfiles(coverage)
			os.Remove(coverage)
		}()
	}
	args := append([]string{"test", "-tags=musl", "-work", "-cpu=2", "-timeout=30s", "-failfast", "-race", "-coverprofile=" + coverage, "-covermode=atomic"})
	args = append(args, file)
	fmt.Printf("exec: go %s\n", strings.Join(args, " "))
	cmd := exec.Command("go", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

func recursiveTest(entry *testSumItem, pkgSum map[string]*testSumItem, incoming map[string]map[string]struct{}, coverage map[string][]*cover.Profile) error {
	if entry.tested {
		return nil
	}
	entry.tested = true
	if entry.info.hasTestFile {
		profiles, err := runTest(string([]rune{'.', os.PathSeparator}) + entry.info.path)
		if err != nil {
			return err
		}
		entry.testTime = time.Now()
		coverage[entry.info.path] = profiles
	}
	for pkg := range incoming[entry.pkg] {
		t := pkgSum[pkg]
		if t != nil {
			err := recursiveTest(t, pkgSum, incoming, coverage)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
