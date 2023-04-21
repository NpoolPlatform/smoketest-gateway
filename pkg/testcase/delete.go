package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

func (h *Handler) DeleteTestCase(ctx context.Context) (*npool.TestCase, error) {
	info, err := h.GetTestCase(ctx)
	if err != nil {
		return nil, err
	}

	_, err = testcasemwcli.DeleteTestCase(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
