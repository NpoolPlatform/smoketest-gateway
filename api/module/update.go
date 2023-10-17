package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-gateway/pkg/module"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) UpdateModule(ctx context.Context, in *npool.UpdateModuleRequest) (*npool.UpdateModuleResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithID(&in.ID, true),
		module1.WithName(in.Name, false),
		module1.WithDescription(in.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateModule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateModuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateModule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateModule",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateModuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateModuleResponse{
		Info: info,
	}, nil
}
