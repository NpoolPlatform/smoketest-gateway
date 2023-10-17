package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-gateway/pkg/module"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) GetModules(ctx context.Context, in *npool.GetModulesRequest) (*npool.GetModulesResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithOffset(in.GetOffset()),
		module1.WithLimit(in.GetLimit()),
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
