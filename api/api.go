package api

import (
	"context"

	smoketest "github.com/NpoolPlatform/message/npool/smoketest/gw/v1"
	"github.com/NpoolPlatform/smoketest-gateway/api/module"
	"github.com/NpoolPlatform/smoketest-gateway/api/testcase"
	"github.com/NpoolPlatform/smoketest-gateway/api/testcase/cond"
	"github.com/NpoolPlatform/smoketest-gateway/api/testplan"
	"github.com/NpoolPlatform/smoketest-gateway/api/testplan/plantestcase"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	smoketest.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	smoketest.RegisterGatewayServer(server, &Server{})
	module.Register(server)
	testcase.Register(server)
	cond.Register(server)
	testplan.Register(server)
	plantestcase.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := smoketest.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := module.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := testcase.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := cond.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := testplan.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := plantestcase.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}
