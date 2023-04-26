package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

//nolint
func (h *Handler) CreateModule(ctx context.Context) (*npool.Module, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid name")
	}

	info, err := cli.CreateModule(
		ctx,
		&npool.ModuleReq{
			Name:        h.Name,
			Description: h.Description,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetModule(ctx)
}
