package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	modulecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) DeleteModule(ctx context.Context) (*npool.Module, error) {
	info, err := h.GetModule(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := modulecli.DeleteModule(ctx, *h.ID); err != nil {
		return nil, err
	}
	return info, nil
}
