// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: task_error.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterTaskErrorServiceImp task_error.proto
func RegisterTaskErrorServiceImp(regester transport.Register, srv TaskErrorServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterTaskErrorServiceHandler(regester, TaskErrorServiceHandler(srv), _ops.HTTP...)
	RegisterTaskErrorServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.dop.taskerror.TaskErrorService",
	)
}

var (
	taskErrorServiceClientType  = reflect.TypeOf((*TaskErrorServiceClient)(nil)).Elem()
	taskErrorServiceServerType  = reflect.TypeOf((*TaskErrorServiceServer)(nil)).Elem()
	taskErrorServiceHandlerType = reflect.TypeOf((*TaskErrorServiceHandler)(nil)).Elem()
)

// TaskErrorServiceClientType .
func TaskErrorServiceClientType() reflect.Type { return taskErrorServiceClientType }

// TaskErrorServiceServerType .
func TaskErrorServiceServerType() reflect.Type { return taskErrorServiceServerType }

// TaskErrorServiceHandlerType .
func TaskErrorServiceHandlerType() reflect.Type { return taskErrorServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		taskErrorServiceClientType,
		// server types
		taskErrorServiceServerType,
		// handler types
		taskErrorServiceHandlerType,
	}
}