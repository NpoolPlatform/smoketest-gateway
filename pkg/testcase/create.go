package testcase

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	if h.ModuleName == nil {
		return fmt.Errorf("invalid module name")
	}
	if h.ApiID == nil {
		return fmt.Errorf("invalid api")
	}
	return nil
}

func (h *Handler) CreateTestCase(ctx context.Context) (*npool.TestCase, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.NewString()
	h.ID = &id

	if _, err := testcasemwcli.CreateTestCase(
		ctx,
		&testcasemgrpb.TestCaseReq{
			ID:          h.ID,
			Name:        h.Name,
			Description: h.Description,
			ModuleName:  h.ModuleName,
			ApiID:       h.ApiID,
			Input:       h.Input,
			InputDesc:   h.InputDesc,
			Expectation: h.Expectation,
		},
	); err != nil {
		return nil, err
	}

	return h.GetTestCase(ctx)
}
