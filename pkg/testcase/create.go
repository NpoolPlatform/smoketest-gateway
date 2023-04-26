package testcase

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
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
	return nil
}

func (h *Handler) CreateTestCase(ctx context.Context) (*npool.TestCase, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	info, err := cli.CreateTestCase(
		ctx,
		&pb.TestCaseReq{
			Name:        h.Name,
			Description: h.Description,
			ModuleName:  h.ModuleName,
			ApiID:       h.ApiID,
			Input:       h.Input,
			InputDesc:   h.InputDesc,
			Expectation: h.Expectation,
			TestCaseType: h.TestCaseType,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetTestCase(ctx)
}
