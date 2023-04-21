package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) DeleteModule(ctx context.Context) (*npool.Module, error) {
	if _, err := modulecli.DeleteModule(ctx, *h.ID); err != nil {
		return nil, err
	}
	return h.GetModule(ctx)
}
