package testcase

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	testcase.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	testcase.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return testcase.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
