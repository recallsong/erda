// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: legacy_consumer.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/consumer/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.hepa.consumer",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                     = reflect.TypeOf((*Client)(nil)).Elem()
	legacyConsumerServiceClientType = reflect.TypeOf((*pb.LegacyConsumerServiceClient)(nil)).Elem()
	legacyConsumerServiceServerType = reflect.TypeOf((*pb.LegacyConsumerServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.hepa.consumer-client":
		return p.client
	case "erda.core.hepa.consumer.LegacyConsumerService":
		return &legacyConsumerServiceWrapper{client: p.client.LegacyConsumerService(), opts: opts}
	case "erda.core.hepa.consumer.LegacyConsumerService.client":
		return p.client.LegacyConsumerService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case legacyConsumerServiceClientType:
		return p.client.LegacyConsumerService()
	case legacyConsumerServiceServerType:
		return &legacyConsumerServiceWrapper{client: p.client.LegacyConsumerService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.consumer-client", &servicehub.Spec{
		Services: []string{
			"erda.core.hepa.consumer.LegacyConsumerService",
			"erda.core.hepa.consumer-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			legacyConsumerServiceClientType,
			// server types
			legacyConsumerServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
