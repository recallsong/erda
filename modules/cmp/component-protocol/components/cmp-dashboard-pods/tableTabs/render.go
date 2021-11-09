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

package tableTabs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

func (tableTabs *TableTabs) Render(ctx context.Context, c *cptype.Component, s cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	if err := tableTabs.GenComponentState(c); err != nil {
		return fmt.Errorf("failed to gen tableTabs component state, %v", err)
	}
	if event.Operation == cptype.InitializeOperation {
		tableTabs.State.ActiveKey = "cpu"
	}
	tableTabs.Props.TabMenu = []TabMenu{
		{
			Key:  "cpu",
			Name: cputil.I18n(ctx, "cpu-analysis"),
		},
		{
			Key:  "mem",
			Name: cputil.I18n(ctx, "mem-analysis"),
		},
	}
	tableTabs.Operations = Operations{
		OnChange: OnChange{
			Key:    "changeTab",
			Reload: true,
		},
	}
	tableTabs.Transfer(c)
	return nil
}

func (tableTabs *TableTabs) GenComponentState(component *cptype.Component) error {
	if component == nil || component.State == nil {
		return nil
	}
	var state State
	data, err := json.Marshal(component.State)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &state); err != nil {
		return err
	}
	tableTabs.State = state
	return nil
}

func (tableTabs *TableTabs) Transfer(c *cptype.Component) {
	c.Props = tableTabs.Props
	c.State = map[string]interface{}{
		"activeKey": tableTabs.State.ActiveKey,
	}
	c.Operations = map[string]interface{}{
		"onChange": tableTabs.Operations.OnChange,
	}
}

func init() {
	base.InitProviderWithCreator("cmp-dashboard-pods", "tableTabs", func() servicehub.Provider {
		return &TableTabs{}
	})
}
