package planrelatedtestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/planrelatedtestcase"
	planrelatedtestcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/planrelatedtestcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

//nolint
func (s *Server) CreatePlanRelatedTestCase(ctx context.Context, in *npool.CreatePlanRelatedTestCaseRequest) (*npool.CreatePlanRelatedTestCaseResponse, error) {
	handler, err := planrelatedtestcase1.NewHandler(
		ctx,
		planrelatedtestcase1.WithTestPlanID(&in.TestPlanID),
		planrelatedtestcase1.WithTestCaseID(&in.TestCaseID),
		planrelatedtestcase1.WithTestUserID(in.TestUserID),
		planrelatedtestcase1.WithTestCaseOutput(in.TestCaseOutput),
		planrelatedtestcase1.WithTestCaseResult(&in.TestCaseResult),
		planrelatedtestcase1.WithIndex(in.Index),
		planrelatedtestcase1.WithRunDuration(in.RunDuration),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePlanRelatedTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePlanRelatedTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePlanRelatedTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePlanRelatedTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePlanRelatedTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePlanRelatedTestCaseResponse{
		Info: info,
	}, nil
}
