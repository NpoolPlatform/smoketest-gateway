package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateTestPlan(ctx context.Context, in *npool.CreateTestPlanRequest) (*npool.CreateTestPlanResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithEntID(in.EntID, false),
		testplan1.WithName(&in.Name, true),
		testplan1.WithCreatedBy(&in.CreatedBy, true),
		testplan1.WithExecutor(in.Executor, false),
		testplan1.WithDeadline(in.Deadline, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTestPlanResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateTestPlan(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTestPlanResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTestPlanResponse{
		Info: info,
	}, nil
}
