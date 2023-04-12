package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) GetTestCases(ctx context.Context, in *npool.GetTestCasesRequest) (*npool.GetTestCasesResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithConds(
			&testcasemgrpb.Conds{
				ModuleID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: in.ModuleID,
				},
			},
			in.GetOffset(),
			in.GetLimit(),
		),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTestCases",
			"In", in,
			"Error", err,
		)
		return &npool.GetTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetTestCases(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTestCases", "err", err)
		return &npool.GetTestCasesResponse{}, err
	}

	return &npool.GetTestCasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
