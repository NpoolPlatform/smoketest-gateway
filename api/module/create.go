package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-gateway/pkg/module"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateModule(ctx context.Context, in *npool.CreateModuleRequest) (*npool.CreateModuleResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithName(&in.Name),
		module1.WithDescription(in.Description),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateModule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateModuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateModule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateModule",
			"In", in,
			"Error", err,
		)
		return &npool.CreateModuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateModuleResponse{
		Info: info,
	}, nil
}
