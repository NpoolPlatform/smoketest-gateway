package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) UpdateTestCase(ctx context.Context, in *npool.UpdateTestCaseRequest) (*npool.UpdateTestCaseResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(&in.ID, true),
		testcase1.WithName(in.Name, false),
		testcase1.WithDescription(in.Description, false),
		testcase1.WithInput(in.Input, false),
		testcase1.WithInputDesc(in.InputDesc, false),
		testcase1.WithExpectation(in.Expectation, false),
		testcase1.WithOutputDesc(in.OutputDesc, false),
		testcase1.WithDeprecated(in.Deprecated, false),
		testcase1.WithTestCaseType(in.TestCaseType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTestCaseResponse{
		Info: info,
	}, nil
}
