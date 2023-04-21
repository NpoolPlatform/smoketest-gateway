package plantestcase

import (
	"context"

	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

//nolint
func (h *Handler) Createplantestcase(ctx context.Context) (*npool.plantestcase, error) {
	info, err := testcasemwcli.Createplantestcase(
		ctx,
		&testcasemwpb.CreateplantestcaseRequest{
			Info: &testcasemwpb.plantestcaseReq{
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

	return h.Getplantestcase(ctx)
}
