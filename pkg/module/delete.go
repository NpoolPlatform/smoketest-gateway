package module

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) DeleteModule(ctx context.Context) (*npool.Module, error) {
	exist, err := cli.ExistModuleConds(ctx, &npool.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}

	info, err := h.GetModule(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := cli.DeleteModule(ctx, *h.ID); err != nil {
		return nil, err
	}
	return info, nil
}
