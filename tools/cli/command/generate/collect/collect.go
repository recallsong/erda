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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	output, err := os.OpenFile("collectCMDs.go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("generating collectCMDs.go fail: %v\n", err)
			os.Remove("collectCMDs.go")
		}
	}()

	if err != nil {
		panic(err)
	}
	pkgs := parseDir("../../cmd/")
	cmds := []string{}
	for _, pkg := range pkgs {
		for fname, f := range pkg.Files {
			cmd := parseFile(fname, f)
			cmds = append(cmds, cmd)
		}
	}
	sort.Strings(cmds)

	io.WriteString(output, "package main\n")
	io.WriteString(output, imports())
	writeCMDs(output, cmds)
	writeCMDNames(output, cmds)
}

func parseDir(path string) map[string]*ast.Package {
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, path, func(info os.FileInfo) bool {
		return true
	}, 0)
	if err != nil {
		panic(err)
	}
	return pkgs
}

func parseFile(fname string, f *ast.File) string {
	for _, decl := range f.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if gen.Tok != token.VAR {
			continue
		}
		for _, spec := range gen.Specs {
			val, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for idx, v := range val.Values {
				e, ok := v.(*ast.CompositeLit)
				if !ok {
					continue
				}
				ident, ok := e.Type.(*ast.Ident)
				if !ok {
					selector, ok := e.Type.(*ast.SelectorExpr)
					if !ok {
						continue
					}
					ident = selector.Sel
				}

				if ident.Name != "Command" {
					continue
				}
				name := val.Names[idx].Name
				if strings.Title(name) != name {
					errStr := fmt.Sprintf("first letter should be uppercase to export it, [%s]", name)
					panic(errStr)
				}
				return name
			}
		}
	}
	errStr := fmt.Sprintf("not found Command in %s", fname)
	panic(errStr)
}

func imports() string {
	return `
import (
        . "github.com/erda-project/erda/tools/cli/command"
        . "github.com/erda-project/erda/tools/cli/cmd"
)
`
}

func writeCMDs(w io.Writer, cmds []string) {
	io.WriteString(w, `
var CMDs = []Command{
`)
	defer io.WriteString(w, "}")
	for _, cmd := range cmds {
		io.WriteString(w, tab(1, cmd))
		io.WriteString(w, ",\n")
	}
}

func writeCMDNames(w io.Writer, cmds []string) {
	io.WriteString(w, `
var CMDNames = []string{
`)
	defer io.WriteString(w, "}")
	for _, cmd := range cmds {
		io.WriteString(w, tab(1, "\""+cmd+"\""))
		io.WriteString(w, ",\n")
	}
}

func tab(n int, content string) string {
	var buf strings.Builder
	for i := 0; i < n; i++ {
		buf.WriteString("	")
	}
	buf.WriteString(content)
	return buf.String()
}
