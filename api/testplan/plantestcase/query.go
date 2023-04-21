package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	prtmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	prt1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan/plantestcase"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

//nolint
func (s *Server) Getplantestcases(ctx context.Context, in *npool.GetplantestcasesRequest) (*npool.GetplantestcasesResponse, error) {
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
			"Getplantestcases",
			"In", in,
			"Error", err,
		)
		return &npool.GetplantestcasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.Getplantestcases(ctx)
	if err != nil {
		logger.Sugar().Errorw("Getplantestcases", "err", err)
		return &npool.GetplantestcasesResponse{}, err
	}

	return &npool.GetplantestcasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
