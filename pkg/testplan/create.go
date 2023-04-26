package testplan

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	if h.CreatedBy == nil {
		return fmt.Errorf("invalid created by")
	}
	return nil
}

func (h *Handler) CreateTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	info, err := testcasemwcli.CreateTestPlan(
		ctx,
		&testplanmgrpb.TestPlanReq{
			ID:        handler.ID,
			Name:      handler.Name,
			CreatedBy: handler.CreatedBy,
			Executor:  handler.Executor,
			Deadline:  handler.Deadline,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetTestPlan(ctx)
}
