package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
)

func (h *Handler) GetConds(ctx context.Context) ([]*npool.Cond, uint32, error) {
	return cli.GetConds(ctx, nil, h.Offset, h.Limit)
}

func (h *Handler) GetCond(ctx context.Context) (*npool.Cond, error) {
	return cli.GetCond(ctx, *h.ID)
}
