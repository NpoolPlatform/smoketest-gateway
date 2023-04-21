package plantestcase

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (h *Handler) DeletePlanTestCase(ctx context.Context) (*mgrpb.PlanTestCase, error) {
	info, err := h.GetPlanTestCase(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := testcasemwcli.DeletePlanTestCase(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
