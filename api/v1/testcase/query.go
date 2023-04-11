package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase"
	testcasemw "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	"github.com/gogo/status"
	"grpc.go4.org/codes"
)

func (s *Server) GetTestCases(ctx context.Context, in *npool.GetTestCasesRequest) (*npool.GetTestCasesResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithConds(
			&testcasemgrpb.Conds{},
			in.GetOffset(),
			in.GetLimit(),
		),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"error", err,
		)
		return &npool.GetTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := testcasemw.GetTestCases()
}
