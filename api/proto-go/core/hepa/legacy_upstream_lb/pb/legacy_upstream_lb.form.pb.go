// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: legacy_upstream_lb.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*UpstreamLb)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TargetOnlineRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TargetOnlineResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TargetOfflineRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TargetOfflineResponse)(nil)

// UpstreamLb implement urlenc.URLValuesUnmarshaler.
func (m *UpstreamLb) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "az":
				m.Az = vals[0]
			case "lbName":
				m.LbName = vals[0]
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "env":
				m.Env = vals[0]
			case "deploymentId":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.DeploymentId = int32(val)
			case "healthcheckPath":
				m.HealthcheckPath = vals[0]
			case "targets":
				m.Targets = vals
			}
		}
	}
	return nil
}

// TargetOnlineRequest implement urlenc.URLValuesUnmarshaler.
func (m *TargetOnlineRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "lb":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
			case "lb.az":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.Az = vals[0]
			case "lb.lbName":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.LbName = vals[0]
			case "lb.orgId":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.OrgId = vals[0]
			case "lb.projectId":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.ProjectId = vals[0]
			case "lb.env":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.Env = vals[0]
			case "lb.deploymentId":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Lb.DeploymentId = int32(val)
			case "lb.healthcheckPath":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.HealthcheckPath = vals[0]
			case "lb.targets":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.Targets = vals
			}
		}
	}
	return nil
}

// TargetOnlineResponse implement urlenc.URLValuesUnmarshaler.
func (m *TargetOnlineResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// TargetOfflineRequest implement urlenc.URLValuesUnmarshaler.
func (m *TargetOfflineRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "lb":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
			case "lb.az":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.Az = vals[0]
			case "lb.lbName":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.LbName = vals[0]
			case "lb.orgId":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.OrgId = vals[0]
			case "lb.projectId":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.ProjectId = vals[0]
			case "lb.env":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.Env = vals[0]
			case "lb.deploymentId":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Lb.DeploymentId = int32(val)
			case "lb.healthcheckPath":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.HealthcheckPath = vals[0]
			case "lb.targets":
				if m.Lb == nil {
					m.Lb = &UpstreamLb{}
				}
				m.Lb.Targets = vals
			}
		}
	}
	return nil
}

// TargetOfflineResponse implement urlenc.URLValuesUnmarshaler.
func (m *TargetOfflineResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}
