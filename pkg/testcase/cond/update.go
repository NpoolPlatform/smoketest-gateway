package cond

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
)

func (h *Handler) UpdateCond(ctx context.Context) (*npool.Cond, error) {
	info, err := cli.GetCondOnly(ctx, &npool.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

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
