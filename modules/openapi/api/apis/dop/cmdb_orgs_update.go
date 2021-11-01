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

package dop

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var CMDB_ORG_UPDATE = apis.ApiSpec{
	Path:         "/api/orgs/<orgID>",
	BackendPath:  "/api/orgs/<orgID>",
	Host:         "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:       "http",
	Method:       "PUT",
	CheckLogin:   true,
	CheckToken:   true,
	IsOpenAPI:    true,
	RequestType:  apistructs.OrgUpdateRequest{},
	ResponseType: apistructs.OrgUpdateRequestBody{},
	Doc:          "summary: 更新组织",
	Audit: func(ctx *spec.AuditContext) error {
		var (
			resp struct{ Data apistructs.OrgDTO }
		)
		if err := ctx.BindResponseData(&resp); err != nil {
			return err
		}

		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.SysScope,
			ScopeID:      1,
			OrgID:        resp.Data.ID,
			TemplateName: apistructs.UpdateOrgTemplate,
			Context: map[string]interface{}{
				"orgName":   resp.Data.Name,
				"contentZH": resp.Data.AuditMessage.MessageZH,
				"contentEN": resp.Data.AuditMessage.MessageEN,
			},
		})
	},
}
