package cond

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
)

func (h *Handler) GetConds(ctx context.Context) ([]*npool.Cond, uint32, error) {
	infos, total, err := cli.GetConds(ctx, nil, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.Cond{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func (h *Handler) GetCond(ctx context.Context) (*npool.Cond, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := cli.GetCond(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return info, nil
}
