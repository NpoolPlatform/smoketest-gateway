package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase/cond"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) DeleteCond(ctx context.Context, in *npool.DeleteCondRequest) (*npool.DeleteCondResponse, error) {
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCond",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCondResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCond(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCond",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCondResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCondResponse{
		Info: info,
	}, nil
}
