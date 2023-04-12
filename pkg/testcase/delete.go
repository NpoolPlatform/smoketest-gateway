package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

type deleteHandler struct {
	*Handler
}

func (h *Handler) DeleteTestCase(ctx context.Context) (*npool.TestCase, error) {
	handler := &deleteHandler{
		Handler: h,
	}

	info, err := handler.DeleteTestCase(ctx)
	if err != nil {
		return nil, err
	}

	_, err = testcasemwcli.DeleteTestCase(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
