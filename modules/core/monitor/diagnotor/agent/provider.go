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

package diagnotor

import (
	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/pkg/transport"
	"github.com/erda-project/erda-proto-go/core/monitor/diagnotor/pb"
	"github.com/erda-project/erda/pkg/common/apis"
)

type config struct {
}

// +provider
type provider struct {
	Cfg                   *config
	Log                   logs.Logger
	Register              transport.Register
	diagnotorAgentService *diagnotorAgentService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.diagnotorAgentService = &diagnotorAgentService{p: p}
	if p.Register != nil {
		pb.RegisterDiagnotorAgentServiceImp(p.Register, p.diagnotorAgentService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.monitor.diagnotor.DiagnotorAgentService" || ctx.Type() == pb.DiagnotorAgentServiceServerType() || ctx.Type() == pb.DiagnotorAgentServiceHandlerType():
		return p.diagnotorAgentService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.diagnotor", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
