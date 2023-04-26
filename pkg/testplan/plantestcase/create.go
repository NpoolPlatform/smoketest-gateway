package plantestcase

import (
	"context"

	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (*pb.PlanTestCase, error) {
	info, err := cli.CreatePlanTestCase(
		ctx,
		&pb.PlanTestCaseReq{
			ID:             h.ID,
			TestPlanID:     h.TestPlanID,
			TestCaseID:     h.TestCaseID,
			TestUserID:     h.TestUserID,
			TestCaseOutput: h.TestCaseOutput,
			Result:         h.Result,
			Index:          h.Index,
			RunDuration:    h.RunDuration,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetPlanTestCase(ctx)
}
