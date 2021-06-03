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

package services

import (
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda/modules/openapi-ng"
)

// +provider
type provider struct {
	Router openapi.Interface `autowired:"openapi-ng"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	RegisterAPIs(p.Router.AddAPI)
	RegisterOldAPIs(p.Router.AddAPI)
	return nil
}

func init() {
	servicehub.Register("openapi-services", &servicehub.Spec{
		Services: []string{"openapi-services"},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
