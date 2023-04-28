package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) DeleteTestPlan(ctx context.Context, in *npool.DeleteTestPlanRequest) (*npool.DeleteTestPlanResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTestPlanResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteTestPlan(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTestPlanResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTestPlanResponse{
		Info: info,
	}, nil
}
