package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

func (h *Handler) CreateTestCase(ctx context.Context) (*npool.TestCase, error) {
	info, err := testcasemwcli.CreateTestCase(ctx, &testcasemwpb.TestCaseReq{
		EntID:         h.EntID,
		Name:          h.Name,
		Description:   h.Description,
		ModuleID:      h.ModuleID,
		ApiID:         h.ApiID,
		Input:         h.Input,
		InputDesc:     h.InputDesc,
		Expectation:   h.Expectation,
		OutputDesc:    h.OutputDesc,
		TestCaseType:  h.TestCaseType,
		TestCaseClass: h.TestCaseClass,
	})
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID
	return h.GetTestCase(ctx)
}
