package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan/plantestcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) DeletePlanTestCase(ctx context.Context, in *npool.DeletePlanTestCaseRequest) (*npool.DeletePlanTestCaseResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithID(&in.ID, true),
		plantestcase1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeletePlanTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePlanTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeletePlanTestCaseResponse{
		Info: info,
	}, nil
}
