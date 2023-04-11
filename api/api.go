package api

import (
	"context"

	smoketest "github.com/NpoolPlatform/message/npool/smoketest/gw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	smoketest.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	smoketest.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := smoketest.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
