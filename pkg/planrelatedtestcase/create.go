package planrelatedtestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/planrelatedtestcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/planrelatedtestcase"
)

//nolint
func (h *Handler) CreatePlanRelatedTestCase(ctx context.Context) (*npool.PlanRelatedTestCase, error) {
	info, err := testcasemwcli.CreatePlanRelatedTestCase(
		ctx,
		&testcasemwpb.CreatePlanRelatedTestCaseRequest{
			Info: &testcasemwpb.PlanRelatedTestCaseReq{
				TestPlanID:     h.TestPlanID,
				TestCaseID:     h.TestCaseID,
				TestUserID:     h.TestUserID,
				TestCaseOutput: h.TestCaseOutput,
				TestCaseResult: h.TestCaseResult,
				Index:          h.Index,
				RunDuration:    h.RunDuration,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID

	return h.GetPlanRelatedTestCase(ctx)
}
