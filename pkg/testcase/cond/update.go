package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
)

func (h *Handler) UpdateCond(ctx context.Context) (*npool.Cond, error) {
	if _, err := cli.UpdateCond(ctx, &npool.CondReq{
		ID:          h.ID,
		CondType:    h.CondType,
		ArgumentMap: h.ArgumentMap,
		Index:       h.Index,
	}); err != nil {
		return nil, err
	}

	return h.GetCond(ctx)
}
