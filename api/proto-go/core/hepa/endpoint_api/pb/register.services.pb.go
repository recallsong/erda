// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: endpoint_api.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterEndpointApiServiceImp endpoint_api.proto
func RegisterEndpointApiServiceImp(regester transport.Register, srv EndpointApiServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterEndpointApiServiceHandler(regester, EndpointApiServiceHandler(srv), _ops.HTTP...)
	RegisterEndpointApiServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.hepa.endpoint_api.EndpointApiService",
	)
}

var (
	endpointApiServiceClientType  = reflect.TypeOf((*EndpointApiServiceClient)(nil)).Elem()
	endpointApiServiceServerType  = reflect.TypeOf((*EndpointApiServiceServer)(nil)).Elem()
	endpointApiServiceHandlerType = reflect.TypeOf((*EndpointApiServiceHandler)(nil)).Elem()
)

// EndpointApiServiceClientType .
func EndpointApiServiceClientType() reflect.Type { return endpointApiServiceClientType }

// EndpointApiServiceServerType .
func EndpointApiServiceServerType() reflect.Type { return endpointApiServiceServerType }

// EndpointApiServiceHandlerType .
func EndpointApiServiceHandlerType() reflect.Type { return endpointApiServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		endpointApiServiceClientType,
		// server types
		endpointApiServiceServerType,
		// handler types
		endpointApiServiceHandlerType,
	}
}
