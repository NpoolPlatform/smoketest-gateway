package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
	"github.com/google/uuid"
)

func (h *Handler) CreateModule(ctx context.Context) (*npool.Module, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid name")
	}

	id := uuid.NewString()
	h.ID = &id

	if _, err := modulecli.CreateModule(
		ctx,
		&npool.ModuleReq{
			ID:          h.ID,
			Name:        h.Name,
			Description: h.Description,
		},
	); err != nil {
		return nil, err
	}

	return h.GetModule(ctx)
}
