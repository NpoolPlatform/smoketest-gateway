package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (*npool.PlanTestCase, error) {
	info, err := cli.CreatePlanTestCase(ctx, &plantestcasemwpb.PlanTestCaseReq{
		EntID:      h.EntID,
		TestPlanID: h.TestPlanID,
		TestCaseID: h.TestCaseID,
		TestUserID: h.TestUserID,
		Index:      h.Index,
		Input:      h.Input,
	})
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	return h.GetPlanTestCase(ctx)
}
