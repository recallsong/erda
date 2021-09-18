// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: settings.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsServiceClient interface {
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error)
	PutSettings(ctx context.Context, in *PutSettingsRequest, opts ...grpc.CallOption) (*PutSettingsResponse, error)
	RegisterMonitorConfig(ctx context.Context, in *RegisterMonitorConfigRequest, opts ...grpc.CallOption) (*RegisterMonitorConfigResponse, error)
}

type settingsServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewSettingsServiceClient(cc grpc1.ClientConnInterface) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error) {
	out := new(GetSettingsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.settings.SettingsService/GetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) PutSettings(ctx context.Context, in *PutSettingsRequest, opts ...grpc.CallOption) (*PutSettingsResponse, error) {
	out := new(PutSettingsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.settings.SettingsService/PutSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) RegisterMonitorConfig(ctx context.Context, in *RegisterMonitorConfigRequest, opts ...grpc.CallOption) (*RegisterMonitorConfigResponse, error) {
	out := new(RegisterMonitorConfigResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.settings.SettingsService/RegisterMonitorConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
// All implementations should embed UnimplementedSettingsServiceServer
// for forward compatibility
type SettingsServiceServer interface {
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	PutSettings(context.Context, *PutSettingsRequest) (*PutSettingsResponse, error)
	RegisterMonitorConfig(context.Context, *RegisterMonitorConfigRequest) (*RegisterMonitorConfigResponse, error)
}

// UnimplementedSettingsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSettingsServiceServer struct {
}

func (*UnimplementedSettingsServiceServer) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (*UnimplementedSettingsServiceServer) PutSettings(context.Context, *PutSettingsRequest) (*PutSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSettings not implemented")
}
func (*UnimplementedSettingsServiceServer) RegisterMonitorConfig(context.Context, *RegisterMonitorConfigRequest) (*RegisterMonitorConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMonitorConfig not implemented")
}

func RegisterSettingsServiceServer(s grpc1.ServiceRegistrar, srv SettingsServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_SettingsService_serviceDesc(srv, opts...), srv)
}

var _SettingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.monitor.settings.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "settings.proto",
}

func _get_SettingsService_serviceDesc(srv SettingsServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_SettingsService_GetSettings_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetSettings(ctx, req.(*GetSettingsRequest))
	}
	var _SettingsService_GetSettings_info transport.ServiceInfo
	if h.Interceptor != nil {
		_SettingsService_GetSettings_info = transport.NewServiceInfo("erda.core.monitor.settings.SettingsService", "GetSettings", srv)
		_SettingsService_GetSettings_Handler = h.Interceptor(_SettingsService_GetSettings_Handler)
	}

	_SettingsService_PutSettings_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.PutSettings(ctx, req.(*PutSettingsRequest))
	}
	var _SettingsService_PutSettings_info transport.ServiceInfo
	if h.Interceptor != nil {
		_SettingsService_PutSettings_info = transport.NewServiceInfo("erda.core.monitor.settings.SettingsService", "PutSettings", srv)
		_SettingsService_PutSettings_Handler = h.Interceptor(_SettingsService_PutSettings_Handler)
	}

	_SettingsService_RegisterMonitorConfig_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RegisterMonitorConfig(ctx, req.(*RegisterMonitorConfigRequest))
	}
	var _SettingsService_RegisterMonitorConfig_info transport.ServiceInfo
	if h.Interceptor != nil {
		_SettingsService_RegisterMonitorConfig_info = transport.NewServiceInfo("erda.core.monitor.settings.SettingsService", "RegisterMonitorConfig", srv)
		_SettingsService_RegisterMonitorConfig_Handler = h.Interceptor(_SettingsService_RegisterMonitorConfig_Handler)
	}

	var serviceDesc = _SettingsService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetSettings",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetSettingsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(SettingsServiceServer).GetSettings(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _SettingsService_GetSettings_info)
				}
				if interceptor == nil {
					return _SettingsService_GetSettings_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.settings.SettingsService/GetSettings",
				}
				return interceptor(ctx, in, info, _SettingsService_GetSettings_Handler)
			},
		},
		{
			MethodName: "PutSettings",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PutSettingsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(SettingsServiceServer).PutSettings(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _SettingsService_PutSettings_info)
				}
				if interceptor == nil {
					return _SettingsService_PutSettings_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.settings.SettingsService/PutSettings",
				}
				return interceptor(ctx, in, info, _SettingsService_PutSettings_Handler)
			},
		},
		{
			MethodName: "RegisterMonitorConfig",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RegisterMonitorConfigRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(SettingsServiceServer).RegisterMonitorConfig(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _SettingsService_RegisterMonitorConfig_info)
				}
				if interceptor == nil {
					return _SettingsService_RegisterMonitorConfig_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.settings.SettingsService/RegisterMonitorConfig",
				}
				return interceptor(ctx, in, info, _SettingsService_RegisterMonitorConfig_Handler)
			},
		},
	}
	return &serviceDesc
}