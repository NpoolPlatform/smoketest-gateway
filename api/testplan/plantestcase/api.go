package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	plantestcase.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	plantestcase.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return plantestcase.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
