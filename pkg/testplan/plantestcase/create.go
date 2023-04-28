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
			TestPlanID: h.TestPlanID,
			TestCaseID: h.TestCaseID,
			TestUserID: h.TestUserID,
			Index:      h.Index,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetPlanTestCase(ctx)
}
