package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

func (h *Handler) DeleteTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	info, err := h.GetTestPlan(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := testplanmwcli.DeleteTestPlan(ctx, *h.ID); err != nil {
		return nil, err
	}

	return info, nil
}
