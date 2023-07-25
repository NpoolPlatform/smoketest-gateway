package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase/cond"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateCond(ctx context.Context, in *npool.CreateCondRequest) (*npool.CreateCondResponse, error) {
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(in.ID),
		cond1.WithTestCaseID(&in.TestCaseID),
		cond1.WithCondTestCaseID(&in.CondTestCaseID),
		cond1.WithCondType(&in.CondType),
		cond1.WithArgumentMap(in.ArgumentMap),
		cond1.WithIndex(&in.Index),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCond",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCondResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCond(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCond",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCondResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCondResponse{
		Info: info,
	}, nil
}
