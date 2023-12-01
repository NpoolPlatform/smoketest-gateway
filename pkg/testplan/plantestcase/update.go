package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) UpdatePlanTestCase(ctx context.Context) (*npool.PlanTestCase, error) {
	info, err := cli.UpdatePlanTestCase(ctx, &plantestcasemwpb.PlanTestCaseReq{
		ID:          h.ID,
		TestUserID:  h.TestUserID,
		Input:       h.Input,
		Output:      h.Output,
		Result:      h.Result,
		Index:       h.Index,
		RunDuration: h.RunDuration,
		Description: h.Description,
	})
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	return h.GetPlanTestCase(ctx)
}
