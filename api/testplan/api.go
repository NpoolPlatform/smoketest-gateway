package testplan

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	testplan.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	testplan.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return testplan.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
