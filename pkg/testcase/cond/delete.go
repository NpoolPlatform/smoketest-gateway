package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
)

func (h *Handler) DeleteCond(ctx context.Context) (*npool.Cond, error) {
	info, err := h.GetCond(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := cli.DeleteCond(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
