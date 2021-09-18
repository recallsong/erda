// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: menu.proto

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

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error)
	GetSetting(ctx context.Context, in *GetSettingRequest, opts ...grpc.CallOption) (*GetSettingResponse, error)
}

type menuServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewMenuServiceClient(cc grpc1.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error) {
	out := new(GetMenuResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.menu.MenuService/GetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetSetting(ctx context.Context, in *GetSettingRequest, opts ...grpc.CallOption) (*GetSettingResponse, error) {
	out := new(GetSettingResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.menu.MenuService/GetSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations should embed UnimplementedMenuServiceServer
// for forward compatibility
type MenuServiceServer interface {
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error)
}

// UnimplementedMenuServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (*UnimplementedMenuServiceServer) GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (*UnimplementedMenuServiceServer) GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetting not implemented")
}

func RegisterMenuServiceServer(s grpc1.ServiceRegistrar, srv MenuServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_MenuService_serviceDesc(srv, opts...), srv)
}

var _MenuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.msp.menu.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "menu.proto",
}

func _get_MenuService_serviceDesc(srv MenuServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_MenuService_GetMenu_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetMenu(ctx, req.(*GetMenuRequest))
	}
	var _MenuService_GetMenu_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MenuService_GetMenu_info = transport.NewServiceInfo("erda.msp.menu.MenuService", "GetMenu", srv)
		_MenuService_GetMenu_Handler = h.Interceptor(_MenuService_GetMenu_Handler)
	}

	_MenuService_GetSetting_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetSetting(ctx, req.(*GetSettingRequest))
	}
	var _MenuService_GetSetting_info transport.ServiceInfo
	if h.Interceptor != nil {
		_MenuService_GetSetting_info = transport.NewServiceInfo("erda.msp.menu.MenuService", "GetSetting", srv)
		_MenuService_GetSetting_Handler = h.Interceptor(_MenuService_GetSetting_Handler)
	}

	var serviceDesc = _MenuService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetMenu",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetMenuRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MenuServiceServer).GetMenu(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MenuService_GetMenu_info)
				}
				if interceptor == nil {
					return _MenuService_GetMenu_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.menu.MenuService/GetMenu",
				}
				return interceptor(ctx, in, info, _MenuService_GetMenu_Handler)
			},
		},
		{
			MethodName: "GetSetting",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetSettingRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(MenuServiceServer).GetSetting(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _MenuService_GetSetting_info)
				}
				if interceptor == nil {
					return _MenuService_GetSetting_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.menu.MenuService/GetSetting",
				}
				return interceptor(ctx, in, info, _MenuService_GetSetting_Handler)
			},
		},
	}
	return &serviceDesc
}
