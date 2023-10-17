package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplanmwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

func (h *Handler) CreateTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	info, err := testplanmwcli.CreateTestPlan(ctx, &testplanmwpb.TestPlanReq{
		ID:        h.ID,
		Name:      h.Name,
		CreatedBy: h.CreatedBy,
		Executor:  h.Executor,
		Deadline:  h.Deadline,
	})
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetTestPlan(ctx)
}
