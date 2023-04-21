package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func (h *Handler) GetModules(ctx context.Context) ([]*npool.Module, uint32, error) {
	infos, total, err := modulemwcli.GetModules(
		ctx,
		&npool.Conds{},
		h.Offset,
		h.Limit,
	)
	if err != nil {
		logger.Sugar().Errorw("GetModules", "err", err)
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

	info, err := modulemwcli.GetModule(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}
	return info, nil
}
