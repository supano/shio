package transport

import (
	endpointv1 "github.com/digithun/shio/admin_api/pkg/endpoint/v1"
	"google.golang.org/grpc"
)

func New(endpoint endpointv1.ShioAdminAPIServer) *grpc.Server {
	s := grpc.NewServer()

	endpointv1.RegisterShioAdminAPIServer(s, endpoint)
	return s
}
