package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan/plantestcase"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) GetPlanTestCases(ctx context.Context, in *npool.GetPlanTestCasesRequest) (*npool.GetPlanTestCasesResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithTestPlanID(&in.TestPlanID, false),
		plantestcase1.WithOffset(in.GetOffset()),
		plantestcase1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPlanTestCases",
			"In", in,
			"Error", err,
		)
		return &npool.GetPlanTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPlanTestCases(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetPlanTestCases", "err", err)
		return &npool.GetPlanTestCasesResponse{}, err
	}

	return &npool.GetPlanTestCasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
