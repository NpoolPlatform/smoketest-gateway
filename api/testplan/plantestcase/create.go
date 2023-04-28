package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan/plantestcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreatePlanTestCase(ctx context.Context, in *npool.CreatePlanTestCaseRequest) (*npool.CreatePlanTestCaseResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithTestPlanID(&in.TestPlanID),
		plantestcase1.WithTestCaseID(&in.TestCaseID),
		plantestcase1.WithIndex(in.Index),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePlanTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePlanTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePlanTestCaseResponse{
		Info: info,
	}, nil
}
