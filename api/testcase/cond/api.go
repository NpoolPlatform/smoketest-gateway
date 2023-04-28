package cond

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase/cond"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	cond.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	cond.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return cond.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
