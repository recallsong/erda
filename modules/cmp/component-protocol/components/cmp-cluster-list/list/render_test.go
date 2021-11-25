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

package list

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/i18n"
	"github.com/erda-project/erda/apistructs"
)

type NopTranslator struct{}

func (t NopTranslator) Get(lang i18n.LanguageCodes, key, def string) string { return key }

func (t NopTranslator) Text(lang i18n.LanguageCodes, key string) string { return key }

func (t NopTranslator) Sprintf(lang i18n.LanguageCodes, key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}

func TestList_GetExtraContent(t *testing.T) {
	type args struct {
		res *ResData
	}
	tests := []struct {
		name string
		args args
		want ExtraContent
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{res: &ResData{
				CpuUsed:     1,
				CpuTotal:    2,
				MemoryUsed:  3,
				MemoryTotal: 4,
				DiskUsed:    9,
				DiskTotal:   10,
			}},
			want: ExtraContent{
				Type:   "PieChart",
				RowNum: 3,
				ExtraData: []ExtraData{
					{
						Name:  "CPU Rate",
						Value: 50.000,
						Total: 100,
						Color: "green",
						Info: []ExtraDataItem{
							{
								Main: "50.000%",
								Sub:  "Rate",
							}, {
								Main: "1.000core",
								Sub:  "Used",
							}, {
								Main: "2.000core",
								Sub:  "CPU" + "Limit",
							},
						},
					},
					{
						Name:  "Memory Rate",
						Value: 75.000,
						Total: 100,
						Color: "green",
						Info: []ExtraDataItem{
							{
								Main: "75.000%",
								Sub:  "Rate",
							}, {
								Main: "3.000",
								Sub:  "Used",
							}, {
								Main: "4.000",
								Sub:  "Memory" + "Limit",
							},
						},
					},
					{
						Name:  "Disk Rate",
						Value: 90.000,
						Total: 100,
						Color: "green",
						Info: []ExtraDataItem{
							{
								Main: "90.000%",
								Sub:  "Rate",
							}, {
								Main: "9.000",
								Sub:  "Used",
							}, {
								Main: "10.000",
								Sub:  "Disk" + "Limit",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				SDK: &cptype.SDK{Tran: NopTranslator{}},
			}
			if got := l.GetExtraContent(tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExtraContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_GetExtraInfos(t *testing.T) {
	type args struct {
		clusterInfo *ClusterInfoDetail
	}
	tests := []struct {
		name string
		args args
		want []ExtraInfos
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{clusterInfo: &ClusterInfoDetail{}},
			want: []ExtraInfos{
				ExtraInfos{
					Icon:    "management",
					Text:    "-",
					Tooltip: "manage type",
				},
				ExtraInfos{
					Icon:    "create-time",
					Text:    "-",
					Tooltip: "create time",
				},
				ExtraInfos{
					Icon:    "machine",
					Text:    "0",
					Tooltip: "machine count",
				},
				ExtraInfos{
					Icon:    "type",
					Text:    "-",
					Tooltip: "cluster type",
				},
				ExtraInfos{
					Icon:    "version",
					Text:    "-",
					Tooltip: "cluster version",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				SDK: &cptype.SDK{Tran: NopTranslator{}},
			}
			if got := l.GetExtraInfos(tt.args.clusterInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExtraInfos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_SetComponentValue(t *testing.T) {
	type args struct {
		c *cptype.Component
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			args:    args{c: &cptype.Component{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{}
			if err := l.SetComponentValue(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetComponentValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_GetBgImage(t *testing.T) {
	type args struct {
		c apistructs.ClusterInfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{c: apistructs.ClusterInfo{Type: "edas"}},
			want: "edas_cluster_bg",
		},
		{
			name: "2",
			args: args{c: apistructs.ClusterInfo{Type: "k8s"}},
			want: "k8s_cluster_bg",
		},
		{
			name: "3",
			args: args{c: apistructs.ClusterInfo{Type: "dcos"}},
			want: "dcos_cluster_bg",
		},
		{
			name: "4",
			args: args{c: apistructs.ClusterInfo{Type: "ack"}},
			want: "ali_cloud_cluster_bg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{}
			if got := l.GetBgImage(tt.args.c); got != tt.want {
				t.Errorf("GetBgImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
