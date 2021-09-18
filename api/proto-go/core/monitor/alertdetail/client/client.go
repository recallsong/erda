// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: details.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/alertdetail/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// AlertDetailService details.proto
	AlertDetailService() pb.AlertDetailServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		alertDetailService: pb.NewAlertDetailServiceClient(cc),
	}
}

type serviceClients struct {
	alertDetailService pb.AlertDetailServiceClient
}

func (c *serviceClients) AlertDetailService() pb.AlertDetailServiceClient {
	return c.alertDetailService
}

type alertDetailServiceWrapper struct {
	client pb.AlertDetailServiceClient
	opts   []grpc1.CallOption
}

func (s *alertDetailServiceWrapper) QuerySystemPodMetrics(ctx context.Context, req *pb.QuerySystemPodMetricsRequest) (*pb.QuerySystemPodMetricsResponse, error) {
	return s.client.QuerySystemPodMetrics(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}