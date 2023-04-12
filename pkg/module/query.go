package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	modulemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

func GetModules(ctx context.Context, offset, limit int32) ([]*npool.Module, uint32, error) {
	infos, total, err := modulemwcli.GetModules(
		ctx,
		&modulemgrpb.Conds{},
		offset,
		limit,
	)
	if err != nil {
		logger.Sugar().Errorw("GetModules", "err", err)
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.Module{}, 0, nil
	}

	_infos := []*npool.Module{}
	for _, info := range infos {
		row := npool.Module{
			ID:          info.ID,
			Name:        info.Name,
			Description: info.Description,
			CreatedAt:   info.CreatedAt,
		}
		_infos = append(_infos, &row)
	}

	return _infos, total, nil
}

func (handler *Handler) GetModule(ctx context.Context) (*npool.Module, error) {
	if handler.ID == nil {
		return nil, fmt.Errorf("invalid module id")
	}

	info, err := modulemwcli.GetModule(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}

	_info := &npool.Module{
		ID:          info.ID,
		Name:        info.Name,
		Description: info.Description,
		CreatedAt:   info.CreatedAt,
	}
	return _info, nil
}
