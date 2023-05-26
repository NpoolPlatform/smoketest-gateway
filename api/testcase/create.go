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
		testcase1.WithName(&in.ID),
		testcase1.WithName(&in.Name),
		testcase1.WithDescription(in.Description),
		testcase1.WithModuleName(&in.ModuleName),
		testcase1.WithApiID(&in.ApiID),
		testcase1.WithInput(in.Input),
		testcase1.WithInputDesc(in.InputDesc),
		testcase1.WithExpectation(in.Expectation),
		testcase1.WithOutputDesc(in.OutputDesc),
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
