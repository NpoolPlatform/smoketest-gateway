package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan/plantestcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) UpdatePlanTestCase(ctx context.Context, in *npool.UpdatePlanTestCaseRequest) (*npool.UpdatePlanTestCaseResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithID(&in.ID),
		plantestcase1.WithTestUserID(in.TestUserID, in.AppID),
		plantestcase1.WithInput(in.Input),
		plantestcase1.WithOutput(in.Output),
		plantestcase1.WithResult(in.Result),
		plantestcase1.WithDescription(in.Description),
		plantestcase1.WithIndex(in.Index),
		plantestcase1.WithRunDuration(in.RunDuration),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdatePlanTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePlanTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePlanTestCaseResponse{
		Info: info,
	}, nil
}
