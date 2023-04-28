package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-gateway/pkg/testcase/cond"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (s *Server) GetConds(ctx context.Context, in *npool.GetCondsRequest) (*npool.GetCondsResponse, error) {
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithOffset(in.GetOffset()),
		cond1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetConds",
			"In", in,
			"Error", err,
		)
		return &npool.GetCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetConds", "err", err)
		return &npool.GetCondsResponse{}, err
	}

	return &npool.GetCondsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
