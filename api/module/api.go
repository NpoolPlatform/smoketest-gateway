package module

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	module.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	module.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return module.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
