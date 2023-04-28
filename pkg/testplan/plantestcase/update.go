package plantestcase

import (
	"context"

	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) UpdatePlanTestCase(ctx context.Context) (*pb.PlanTestCase, error) {
	info, err := cli.UpdatePlanTestCase(
		ctx,
		&pb.PlanTestCaseReq{
			ID:             h.ID,
			TestUserID:     h.TestUserID,
			TestCaseOutput: h.TestCaseOutput,
			Result:         h.Result,
			Index:          h.Index,
			RunDuration:    h.RunDuration,
			Description:    h.Description,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetPlanTestCase(ctx)
}
