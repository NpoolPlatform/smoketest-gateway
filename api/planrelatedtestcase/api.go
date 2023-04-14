package planrelatedtestcase

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/planrelatedtestcase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	planrelatedtestcase.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	planrelatedtestcase.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return planrelatedtestcase.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
