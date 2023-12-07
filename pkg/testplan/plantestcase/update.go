package plantestcase

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) UpdatePlanTestCase(ctx context.Context) (*npool.PlanTestCase, error) {
	info, err := cli.GetPlanTestCaseOnly(ctx, &plantestcasemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	_, err = cli.UpdatePlanTestCase(ctx, &plantestcasemwpb.PlanTestCaseReq{
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

	return h.GetPlanTestCase(ctx)
}
