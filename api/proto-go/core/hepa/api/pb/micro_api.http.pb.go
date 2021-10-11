// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: micro_api.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ApiServiceHandler is the server API for ApiService service.
type ApiServiceHandler interface {
	// +publish path:"/api/gateway/api"
	// GET /api/gateway/api
	GetApis(context.Context, *GetApisRequest) (*GetApisResponse, error)
	// +publish path:"/api/gateway/api"
	// POST /api/gateway/api
	CreateApi(context.Context, *CreateApiRequest) (*CreateApiResponse, error)
	// +publish path:"/api/gateway/api/{apiId}"
	// PATCH /api/gateway/api/{apiId}
	UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiResponse, error)
	// +publish path:"/api/gateway/api/{apiId}"
	// DELETE /api/gateway/api/{apiId}
	DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiResponse, error)
}

// RegisterApiServiceHandler register ApiServiceHandler to http.Router.
func RegisterApiServiceHandler(r http.Router, srv ApiServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_GetApis := func(method, path string, fn func(context.Context, *GetApisRequest) (*GetApisResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetApisRequest))
		}
		var GetApis_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetApis_info = transport.NewServiceInfo("erda.core.hepa.api.ApiService", "GetApis", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetApis_info)
				}
				r = r.WithContext(ctx)
				var in GetApisRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["apiPath"]; len(vals) > 0 {
					in.ApiPath = vals[0]
				}
				if vals := params["diceApp"]; len(vals) > 0 {
					in.DiceApp = vals[0]
				}
				if vals := params["diceService"]; len(vals) > 0 {
					in.DiceService = vals[0]
				}
				if vals := params["env"]; len(vals) > 0 {
					in.Env = vals[0]
				}
				if vals := params["from"]; len(vals) > 0 {
					in.From = vals[0]
				}
				if vals := params["method"]; len(vals) > 0 {
					in.Method = vals[0]
				}
				if vals := params["needAuth"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 32)
					if err != nil {
						return nil, err
					}
					in.NeedAuth = int32(val)
				}
				if vals := params["netType"]; len(vals) > 0 {
					in.NetType = vals[0]
				}
				if vals := params["orgId"]; len(vals) > 0 {
					in.OrgId = vals[0]
				}
				if vals := params["page"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.Page = val
				}
				if vals := params["projectId"]; len(vals) > 0 {
					in.ProjectId = vals[0]
				}
				if vals := params["registerType"]; len(vals) > 0 {
					in.RegisterType = vals[0]
				}
				if vals := params["runtimeId"]; len(vals) > 0 {
					in.RuntimeId = vals[0]
				}
				if vals := params["size"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.Size = val
				}
				if vals := params["sortField"]; len(vals) > 0 {
					in.SortField = vals[0]
				}
				if vals := params["sortType"]; len(vals) > 0 {
					in.SortType = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateApi := func(method, path string, fn func(context.Context, *CreateApiRequest) (*CreateApiResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateApiRequest))
		}
		var CreateApi_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateApi_info = transport.NewServiceInfo("erda.core.hepa.api.ApiService", "CreateApi", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateApi_info)
				}
				r = r.WithContext(ctx)
				var in CreateApiRequest
				if err := h.Decode(r, &in.ApiRequest); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateApi := func(method, path string, fn func(context.Context, *UpdateApiRequest) (*UpdateApiResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateApiRequest))
		}
		var UpdateApi_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateApi_info = transport.NewServiceInfo("erda.core.hepa.api.ApiService", "UpdateApi", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateApi_info)
				}
				r = r.WithContext(ctx)
				var in UpdateApiRequest
				if err := h.Decode(r, &in.ApiRequest); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "apiId":
							in.ApiId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteApi := func(method, path string, fn func(context.Context, *DeleteApiRequest) (*DeleteApiResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteApiRequest))
		}
		var DeleteApi_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteApi_info = transport.NewServiceInfo("erda.core.hepa.api.ApiService", "DeleteApi", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteApi_info)
				}
				r = r.WithContext(ctx)
				var in DeleteApiRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "apiId":
							in.ApiId = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetApis("GET", "/api/gateway/api", srv.GetApis)
	add_CreateApi("POST", "/api/gateway/api", srv.CreateApi)
	add_UpdateApi("PATCH", "/api/gateway/api/{apiId}", srv.UpdateApi)
	add_DeleteApi("DELETE", "/api/gateway/api/{apiId}", srv.DeleteApi)
}
