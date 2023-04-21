package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan/plantestcase"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

//nolint
func (s *Server) Createplantestcase(ctx context.Context, in *npool.CreateplantestcaseRequest) (*npool.CreateplantestcaseResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithTestPlanID(&in.TestPlanID),
		plantestcase1.WithTestCaseID(&in.TestCaseID),
		plantestcase1.WithTestUserID(in.TestUserID),
		plantestcase1.WithTestCaseOutput(in.TestCaseOutput),
		plantestcase1.WithTestCaseResult(&in.TestCaseResult),
		plantestcase1.WithIndex(in.Index),
		plantestcase1.WithRunDuration(in.RunDuration),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Createplantestcase",
			"In", in,
			"Error", err,
		)
		return &npool.CreateplantestcaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.Createplantestcase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Createplantestcase",
			"In", in,
			"Error", err,
		)
		return &npool.CreateplantestcaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateplantestcaseResponse{
		Info: info,
	}, nil
}
