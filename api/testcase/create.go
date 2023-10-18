package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateTestCase(ctx context.Context, in *npool.CreateTestCaseRequest) (*npool.CreateTestCaseResponse, error) {
	apiMust := false
	if in.TestCaseType == testcasemwpb.TestCaseType_Automatic {
		apiMust = true
	}

	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(in.ID, false),
		testcase1.WithName(&in.Name, true),
		testcase1.WithDescription(in.Description, false),
		testcase1.WithModuleID(&in.ModuleID, true),
		testcase1.WithApiID(in.ApiID, apiMust),
		testcase1.WithInput(in.Input, false),
		testcase1.WithInputDesc(in.InputDesc, false),
		testcase1.WithExpectation(in.Expectation, false),
		testcase1.WithOutputDesc(in.OutputDesc, false),
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
