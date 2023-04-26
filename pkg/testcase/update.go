package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

type updateHandler struct {
	*Handler
}

func (h *Handler) UpdateTestCase(ctx context.Context) (*npool.TestCase, error) {
	_, err := testcasemwcli.UpdateTestCase(ctx, &testcasemgrpb.TestCaseReq{
		ID:           h.ID,
		Name:         h.Name,
		Description:  h.Description,
		Input:        h.Input,
		InputDesc:    h.InputDesc,
		Expectation:  h.Expectation,
		TestCaseType: h.TestCaseType,
		Deprecated:   h.Deprecated,
	},
	)
	if err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler: h,
	}

	info, err := handler.GetTestCase(ctx)
	if err != nil {
		return nil, err
	}

	return info, nil
}
