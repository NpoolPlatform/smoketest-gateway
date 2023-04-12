package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/module"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	if *h.Name == "" {
		return fmt.Errorf("invalid name")
	}
	return nil
}

func (h *Handler) CreateModule(ctx context.Context) (*npool.Module, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	info, err := testcasemwcli.CreateModule(
		ctx,
		&testcasemwpb.CreateModuleRequest{
			Name:        *h.Name,
			Description: h.Description,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID

	return h.GetModule(ctx)
}
