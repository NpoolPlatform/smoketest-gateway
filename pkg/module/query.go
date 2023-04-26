package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) GetModules(ctx context.Context) ([]*npool.Module, uint32, error) {
	infos, total, err := cli.GetModules(
		ctx,
		nil,
		h.Offset,
		h.Limit,
	)
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func (h *Handler) GetModule(ctx context.Context) (*npool.Module, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid module id")
	}

	info, err := cli.GetModule(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
