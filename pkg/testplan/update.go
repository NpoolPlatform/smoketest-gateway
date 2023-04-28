package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

func (h *Handler) UpdateTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	if _, err := cli.UpdateTestPlan(
		ctx,
		&pb.TestPlanReq{
			ID:          h.ID,
			Name:        h.Name,
			Executor:    h.Executor,
			State:       h.State,
			Deadline:    h.Deadline,
			Fails:       h.Fails,
			Skips:       h.Skips,
			Passes:      h.Passes,
			RunDuration: h.RunDuration,
			Result:      h.Result,
		},
	); err != nil {
		return nil, err
	}

	return h.GetTestPlan(ctx)
}
