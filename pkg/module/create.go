package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) CreateModule(ctx context.Context) (*npool.Module, error) {
	info, err := cli.CreateModule(ctx, &npool.ModuleReq{
		Name:        h.Name,
		Description: h.Description,
	})
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetModule(ctx)
}
