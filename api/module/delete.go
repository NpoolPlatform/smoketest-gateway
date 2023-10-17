package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-gateway/pkg/module"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) DeleteModule(ctx context.Context, in *npool.DeleteModuleRequest) (*npool.DeleteModuleResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteModule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteModuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteModule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteModule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteModuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteModuleResponse{
		Info: info,
	}, nil
}
