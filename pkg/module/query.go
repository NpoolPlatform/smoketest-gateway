package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) GetModules(ctx context.Context) ([]*npool.Module, uint32, error) {
	return cli.GetModules(
		ctx,
		nil,
		h.Offset,
		h.Limit,
	)
}

func (h *Handler) GetModule(ctx context.Context) (*npool.Module, error) {
	return cli.GetModule(ctx, *h.EntID)
}
