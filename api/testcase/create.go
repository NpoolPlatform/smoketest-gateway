package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateTestCase(ctx context.Context, in *npool.CreateTestCaseRequest) (*npool.CreateTestCaseResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(in.ID, false),
		testcase1.WithName(&in.Name, true),
		testcase1.WithDescription(in.Description, false),
		testcase1.WithModuleID(&in.ModuleID, true),
		testcase1.WithApiID(&in.ApiID, false),
		testcase1.WithInput(in.Input, true),
		testcase1.WithInputDesc(in.InputDesc, true),
		testcase1.WithExpectation(in.Expectation, true),
		testcase1.WithOutputDesc(in.OutputDesc, true),
		testcase1.WithTestCaseType(&in.TestCaseType, true),
		testcase1.WithTestCaseClass(&in.TestCaseClass, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTestCaseResponse{
		Info: info,
	}, nil
}
