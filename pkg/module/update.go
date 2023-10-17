package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) UpdateModule(ctx context.Context) (*npool.Module, error) {
	if _, err := cli.UpdateModule(ctx, &npool.ModuleReq{
		ID:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}); err != nil {
		return nil, err
	}

	return h.GetModule(ctx)
}
