package filemanager

import (
	"context"
	"fmt"

	"github.com/erda-project/erda-infra/providers/httpserver"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/monitor/common/permission"
	perm "github.com/erda-project/erda/pkg/common/permission"
)

func (p *provider) getScopeByHTTPRequest(ctx httpserver.Context) (string, error) {
	fmt.Println("scope", ctx.Request().URL.Query().Get("scope"))
	switch ctx.Request().URL.Query().Get("scope") {
	case "org":
		return permission.ScopeOrg, nil
	case "app":
		return permission.ScopeApp, nil
	}
	return "", fmt.Errorf("invalid scope")
}

func (p *provider) checkScopeIDByHTTPRequest(ctx httpserver.Context) (string, error) {
	params := ctx.Request().URL.Query()
	setInstance := func(instance *apistructs.InstanceInfoData) *apistructs.InstanceInfoData {
		ctx.SetAttribute(instanceKey, instance)
		return instance
	}
	if params.Get("scope") == "org" {
		return p.checkOrgScopeID(ctx.Param("containerID"), params.Get("hostIP"), permission.ScopeOrg, setInstance)
	}
	return p.checkOrgScopeID(ctx.Param("containerID"), params.Get("hostIP"), permission.ScopeApp, setInstance)
}

func (p *provider) getScopeByRequest(ctx context.Context, req interface{}) (string, error) {
	r, ok := req.(reuqestForPermission)
	if !ok {
		return "", fmt.Errorf("invalid reuqest")
	}
	switch r.GetScope() {
	case "org":
		return perm.ScopeOrg, nil
	case "app":
		return perm.ScopeApp, nil
	}
	return "", fmt.Errorf("invalid scope")
}

func (p *provider) checkScopeID(ctx context.Context, req interface{}) (string, error) {
	r, ok := req.(reuqestForPermission)
	if !ok {
		return "", fmt.Errorf("invalid reuqest")
	}
	setInstance := func(instance *apistructs.InstanceInfoData) *apistructs.InstanceInfoData {
		perm.SetPermissionDataFromContext(ctx, instanceKey, instance)
		return instance
	}
	if r.GetScope() == "org" {
		return p.checkOrgScopeID(r.GetContainerID(), r.GetHostIP(), permission.ScopeOrg, setInstance)
	}
	return p.checkOrgScopeID(r.GetContainerID(), r.GetHostIP(), permission.ScopeOrg, setInstance)
}

type reuqestForPermission interface {
	GetContainerID() string
	GetHostIP() string
	GetScope() string
}

func (p *provider) checkOrgScopeID(containerID, hostIP, scope string, fn func(instance *apistructs.InstanceInfoData) *apistructs.InstanceInfoData) (string, error) {
	resp, err := p.bdl.GetInstanceInfo(apistructs.InstanceInfoRequest{
		ContainerID: containerID,
		HostIP:      hostIP,
		Limit:       1,
	})
	if err != nil {
		return "", fmt.Errorf("failed to GetInstanceInfo: %s", err)
	}
	if !resp.Success {

		return "", fmt.Errorf("failed to GetInstanceInfo: code=%s, msg=%s", resp.Error.Code, resp.Error.Msg)
	}
	if len(resp.Data) <= 0 {
		return "", fmt.Errorf("not found instance %s/%s", hostIP, containerID)
	}
	instance := &resp.Data[0]
	if fn != nil {
		instance = fn(instance)
	}
	switch scope {
	case "org":
		return instance.OrgID, nil
	case "app":
		return instance.ApplicationID, nil
	default:
		return "", fmt.Errorf("invalid scope")
	}
}
