package plantestcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	mwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) GetPlanTestCases(ctx context.Context) ([]*mgrpb.PlanTestCase, uint32, error) {
	infos, total, err := mwcli.GetPlanTestCases(ctx, &mgrpb.Conds{
		TestPlanID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: *h.TestPlanID,
		},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*mgrpb.PlanTestCase{}, 0, nil
	}
	return infos, total, nil
}

func (h *Handler) GetPlanTestCase(ctx context.Context) (*mgrpb.PlanTestCase, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := mwcli.GetPlanTestCase(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
