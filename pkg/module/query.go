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
		&npool.Conds{},
		h.Offset,
		h.Limit,
	)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.Module{}, 0, nil
	}

	return infos, total, nil
}

func (handler *Handler) GetModule(ctx context.Context) (*npool.Module, error) {
	if handler.ID == nil {
		return nil, fmt.Errorf("invalid module id")
	}

	info, err := cli.GetModule(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
