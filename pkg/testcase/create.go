package testcase

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	if h.ModuleName == nil {
		return fmt.Errorf("invalid module name")
	}
	if h.ApiID == nil {
		return fmt.Errorf("invalid api")
	}
	if h.Arguments == nil {
		return fmt.Errorf("invalid arguments")
	}
	if h.ExpectationResult == nil {
		return fmt.Errorf("invalid expectation")
	}

	return nil
}

//nolint
func (h *Handler) CreateTestCase(ctx context.Context) (*npool.TestCase, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	info, err := testcasemwcli.CreateTestCase(
		ctx,
		&testcasemwpb.CreateTestCaseReq{
			Name:              *h.Name,
			Description:       h.Description,
			ModuleName:        h.ModuleName,
			ApiID:             *h.ApiID,
			Arguments:         *h.Arguments,
			ExpectationResult: *h.ExpectationResult,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID

	return h.GetTestCase(ctx)
}
