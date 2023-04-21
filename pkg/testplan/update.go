package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

func (h *Handler) UpdateTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	if _, err := testcasemwcli.UpdateTestPlan(
		ctx,
		&testplanmgrpb.TestPlanReq{
			ID:       h.ID,
			Name:     h.Name,
			Executor: h.Executor,
			State:    h.State,
			Deadline: h.Deadline,
		},
	); err != nil {
		return nil, err
	}

	return h.GetTestPlan(ctx)
}
