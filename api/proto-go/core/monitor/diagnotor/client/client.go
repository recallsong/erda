// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: diagnotor.proto, diagnotor_agent.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/diagnotor/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// DiagnotorService diagnotor.proto
	DiagnotorService() pb.DiagnotorServiceClient
	// DiagnotorAgentService diagnotor_agent.proto
	DiagnotorAgentService() pb.DiagnotorAgentServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		diagnotorService:      pb.NewDiagnotorServiceClient(cc),
		diagnotorAgentService: pb.NewDiagnotorAgentServiceClient(cc),
	}
}

type serviceClients struct {
	diagnotorService      pb.DiagnotorServiceClient
	diagnotorAgentService pb.DiagnotorAgentServiceClient
}

func (c *serviceClients) DiagnotorService() pb.DiagnotorServiceClient {
	return c.diagnotorService
}

func (c *serviceClients) DiagnotorAgentService() pb.DiagnotorAgentServiceClient {
	return c.diagnotorAgentService
}

type diagnotorServiceWrapper struct {
	client pb.DiagnotorServiceClient
	opts   []grpc1.CallOption
}

func (s *diagnotorServiceWrapper) StartDiagnosis(ctx context.Context, req *pb.StartDiagnosisRequest) (*pb.StartDiagnosisResponse, error) {
	return s.client.StartDiagnosis(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *diagnotorServiceWrapper) QueryDiagnosisStatus(ctx context.Context, req *pb.QueryDiagnosisStatusRequest) (*pb.QueryDiagnosisStatusResponse, error) {
	return s.client.QueryDiagnosisStatus(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *diagnotorServiceWrapper) ListProcesses(ctx context.Context, req *pb.ListProcessesRequest) (*pb.ListProcessesResponse, error) {
	return s.client.ListProcesses(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

type diagnotorAgentServiceWrapper struct {
	client pb.DiagnotorAgentServiceClient
	opts   []grpc1.CallOption
}

func (s *diagnotorAgentServiceWrapper) ListTargetProcesses(ctx context.Context, req *pb.ListTargetProcessesRequest) (*pb.ListTargetProcessesResponse, error) {
	return s.client.ListTargetProcesses(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
