package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	modulemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-gateway/pkg/module"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) GetModules(ctx context.Context, in *npool.GetModulesRequest) (*npool.GetModulesResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithConds(
			&modulemgrpb.Conds{},
			in.GetOffset(),
			in.GetLimit(),
		),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetModules",
			"In", in,
			"Error", err,
		)
		return &npool.GetModulesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetModules(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetModules", "err", err)
		return &npool.GetModulesResponse{}, err
	}

	return &npool.GetModulesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
