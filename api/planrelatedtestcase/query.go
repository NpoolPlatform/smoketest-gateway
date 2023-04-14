package planrelatedtestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/planrelatedtestcase"
	prtmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	prt1 "github.com/NpoolPlatform/smoketest-gateway/pkg/planrelatedtestcase"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

//nolint
func (s *Server) GetPlanRelatedTestCases(ctx context.Context, in *npool.GetPlanRelatedTestCasesRequest) (*npool.GetPlanRelatedTestCasesResponse, error) {
	handler, err := prt1.NewHandler(
		ctx,
		prt1.WithConds(
			&prtmgrpb.Conds{},
			in.GetOffset(),
			in.GetLimit(),
		),
		prt1.WithTestPlanID(&in.TestPlanID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPlanRelatedTestCases",
			"In", in,
			"Error", err,
		)
		return &npool.GetPlanRelatedTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPlanRelatedTestCases(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetPlanRelatedTestCases", "err", err)
		return &npool.GetPlanRelatedTestCasesResponse{}, err
	}

	return &npool.GetPlanRelatedTestCasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
