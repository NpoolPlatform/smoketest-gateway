package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) DeleteTestCase(ctx context.Context, in *npool.DeleteTestCaseRequest) (*npool.DeleteTestCaseResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTestCaseResponse{
		Info: info,
	}, nil
}
