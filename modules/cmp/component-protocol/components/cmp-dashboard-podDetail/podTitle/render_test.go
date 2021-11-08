// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package podTitle

import (
	"testing"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/cmp/component-protocol/cputil"
)

func TestPodTitle_GenComponentState(t *testing.T) {
	c := &cptype.Component{State: map[string]interface{}{
		"podId": "testPodID",
	}}
	podTitle := &PodTitle{}
	if err := podTitle.GenComponentState(c); err != nil {
		t.Fatal(err)
	}

	ok, err := cputil.IsJsonEqual(c.State, podTitle.State)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Errorf("test failed, json is not equal")
	}
}

func TestPodTitle_Transfer(t *testing.T) {
	podTitle := &PodTitle{
		Props: Props{
			Title: "testTitle",
		},
		State: State{
			PodID: "testPodId",
		},
	}

	c := &cptype.Component{}
	podTitle.Transfer(c)

	ok, err := cputil.IsJsonEqual(c, podTitle)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Errorf("test failed, json is not equal")
	}
}
