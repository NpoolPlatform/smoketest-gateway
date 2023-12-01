package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase/cond"
)

func (h *Handler) CreateCond(ctx context.Context) (*npool.Cond, error) {
	info, err := cli.CreateCond(ctx, &npool.CondReq{
		EntID:          h.EntID,
		TestCaseID:     h.TestCaseID,
		CondTestCaseID: h.CondTestCaseID,
		CondType:       h.CondType,
		ArgumentMap:    h.ArgumentMap,
		Index:          h.Index,
	})
	if err != nil {
		return nil, err
	}

	h.EntID = &info.EntID
	h.ID = &info.ID
	return h.GetCond(ctx)
}
