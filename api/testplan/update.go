package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) UpdateTestPlan(ctx context.Context, in *npool.UpdateTestPlanRequest) (*npool.UpdateTestPlanResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(&in.ID),
		testplan1.WithName(in.Name),
		testplan1.WithExecutor(in.Executor),
		testplan1.WithState(in.State),
		testplan1.WithDeadline(in.Deadline),
		testplan1.WithFails(in.Fails),
		testplan1.WithSkips(in.Skips),
		testplan1.WithFails(in.Fails),
		testplan1.WithResult(in.Result),
		testplan1.WithRunDuration(in.RunDuration),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTestPlanResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateTestPlan(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTestPlanResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTestPlanResponse{
		Info: info,
	}, nil
}
