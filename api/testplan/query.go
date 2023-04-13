package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testplan"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) GetTestPlans(ctx context.Context, in *npool.GetTestPlansRequest) (*npool.GetTestPlansResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithConds(
			&testplanmgrpb.Conds{},
			in.GetOffset(),
			in.GetLimit(),
		),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTestPlans",
			"In", in,
			"Error", err,
		)
		return &npool.GetTestPlansResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetTestPlans(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTestPlans", "err", err)
		return &npool.GetTestPlansResponse{}, err
	}

	return &npool.GetTestPlansResponse{
		Infos: infos,
		Total: total,
	}, nil
}
