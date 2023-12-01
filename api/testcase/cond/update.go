package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase/cond"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) UpdateCond(ctx context.Context, in *npool.UpdateCondRequest) (*npool.UpdateCondResponse, error) {
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(&in.ID, true),
		cond1.WithEntID(&in.EntID, true),
		cond1.WithCondType(in.CondType, false),
		cond1.WithArgumentMap(in.ArgumentMap, false),
		cond1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCond",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCondResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCond(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCond",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCondResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCondResponse{
		Info: info,
	}, nil
}
