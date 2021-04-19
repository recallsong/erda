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

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/strutil"
	"github.com/erda-project/erda/tools/cli/command"
	"github.com/erda-project/erda/tools/cli/format"
)

func GetApplicationDetail(ctx *command.Context, application string, project string) (
	apistructs.ApplicationFetchResponse, error) {
	var (
		resp  apistructs.ApplicationFetchResponse
		b     bytes.Buffer
		appID string
	)

	if application == "" {
		return apistructs.ApplicationFetchResponse{}, fmt.Errorf(format.FormatErrMsg("get application detail",
			"missing required parameter application", false))
	}

	appID = application

	if project != "" {
		appFlag := false
		appList, err := GetApplicationList(ctx, project)
		if err != nil {
			return apistructs.ApplicationFetchResponse{}, err
		}

		for i := range appList {
			if appList[i].Name == application {
				appID = strconv.FormatUint(appList[i].ID, 10)
				appFlag = true
			}
		}

		if !appFlag {
			return apistructs.ApplicationFetchResponse{}, fmt.Errorf(format.FormatErrMsg("get application detail",
				"failed to get app from apps list", false))
		}
	}

	response, err := ctx.Get().Path(strutil.Concat("/api/applications/", appID)).Do().Body(&b)
	if err != nil {
		return apistructs.ApplicationFetchResponse{}, fmt.Errorf(format.FormatErrMsg(
			"get application detail", "failed to request ("+err.Error()+")", false))
	}

	if !response.IsOK() {
		return apistructs.ApplicationFetchResponse{}, fmt.Errorf(format.FormatErrMsg("get application detail",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				response.StatusCode(), response.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &resp); err != nil {
		return apistructs.ApplicationFetchResponse{}, fmt.Errorf(format.FormatErrMsg("get application detail",
			fmt.Sprintf("failed to unmarshal application detail response ("+err.Error()+")"), false))
	}

	if !resp.Success {
		return apistructs.ApplicationFetchResponse{}, fmt.Errorf(format.FormatErrMsg("get application detail",
			fmt.Sprintf("failed to request, error code: %s, error message: %s",
				resp.Error.Code, resp.Error.Msg), false))
	}

	return resp, nil
}

func GetApplicationList(ctx *command.Context, project string) ([]apistructs.ApplicationDTO, error) {
	var resp apistructs.ApplicationListResponse
	var b bytes.Buffer

	response, err := ctx.Get().Path("/api/applications").Param("projectId", project).
		Param("pageSize", "200").Do().Body(&b)
	if err != nil {
		return nil, fmt.Errorf(
			format.FormatErrMsg("list", "failed to request ("+err.Error()+")", false))
	}

	if !response.IsOK() {
		return nil, fmt.Errorf(format.FormatErrMsg("list",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				response.StatusCode(), response.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf(format.FormatErrMsg("list",
			fmt.Sprintf("failed to unmarshal application list response ("+err.Error()+")"), false))
	}

	if !resp.Success {
		return nil, fmt.Errorf(
			format.FormatErrMsg("list",
				fmt.Sprintf("failed to request, error code: %s, error message: %s",
					resp.Error.Code, resp.Error.Msg), false))
	}

	if resp.Data.Total == 0 {
		fmt.Printf(format.FormatErrMsg("list", "no applications created\n", false))
		return nil, nil
	}

	return resp.Data.List, nil
}
