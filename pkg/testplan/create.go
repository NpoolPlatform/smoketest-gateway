package testplan

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	if h.OwnerID == nil {
		return fmt.Errorf("invalid owner id")
	}
	return nil
}

//nolint
func (h *Handler) CreateTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	info, err := testcasemwcli.CreateTestPlan(
		ctx,
		&testcasemwpb.CreateTestPlanRequest{
			Info: &testcasemwpb.TestPlanReq{
				Name:              handler.Name,
				OwnerID:           handler.OwnerID,
				ResponsibleUserID: handler.ResponsibleUserID,
				Deadline:          handler.Deadline,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID

	return h.GetTestPlan(ctx)
}
