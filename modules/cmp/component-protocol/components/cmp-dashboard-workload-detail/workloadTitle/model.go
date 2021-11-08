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

package workloadTitle

import (
	"context"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/cmp/cmp_interface"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

type ComponentWorkloadTitle struct {
	base.DefaultProvider

	sdk    *cptype.SDK
	ctx    context.Context
	server cmp_interface.SteveServer
	Type   string `json:"type,omitempty"`
	Props  Props  `json:"props,omitempty"`
	State  State  `json:"state,omitempty"`
}

type Props struct {
	Title string `json:"title,omitempty"`
}

type State struct {
	ClusterName string `json:"clusterName,omitempty"`
	PodID       string `json:"podId,omitempty"`
	WorkloadID  string `json:"workloadId,omitempty"`
}
