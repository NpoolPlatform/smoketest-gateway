package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
	testcasecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
	"github.com/google/uuid"
)

func (h *Handler) CreateCond(ctx context.Context) (*npool.Cond, error) {
	id := uuid.NewString()
	h.ID = &id

	if _, err := testcasecli.CreateCond(
		ctx,
		&npool.CondReq{
			ID:             h.ID,
			TestCaseID:     h.TestCaseID,
			CondTestCaseID: h.CondTestCaseID,
			CondType:       h.CondType,
			ArgumentMap:    h.ArgumentMap,
			Index:          h.Index,
		},
	); err != nil {
		return nil, err
	}

	return h.GetCond(ctx)
}
