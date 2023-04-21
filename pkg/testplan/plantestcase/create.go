package plantestcase

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
	"github.com/google/uuid"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (*mgrpb.PlanTestCase, error) {
	id := uuid.NewString()
	h.ID = &id

	if _, err := testcasemwcli.CreatePlanTestCase(
		ctx,
		&mgrpb.PlanTestCaseReq{
			ID:             h.ID,
			TestPlanID:     h.TestPlanID,
			TestCaseID:     h.TestCaseID,
			TestUserID:     h.TestUserID,
			TestCaseOutput: h.TestCaseOutput,
			Result:         h.Result,
			Index:          h.Index,
			RunDuration:    h.RunDuration,
		},
	); err != nil {
		return nil, err
	}

	return h.GetPlanTestCase(ctx)
}
