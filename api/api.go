package api

import (
	"context"

	testcasegw "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	testcasegw.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	testcasegw.RegisterGatewayServer(server, &Server{})
}
func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := testcasegw.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
