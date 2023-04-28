package testplan

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	return nil
}

func (h *Handler) CreateTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	info, err := cli.CreateTestPlan(
		ctx,
		&pb.TestPlanReq{
			Name:      handler.Name,
			CreatedBy: handler.CreatedBy,
			Executor:  handler.Executor,
			Deadline:  handler.Deadline,
		},
	)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	return h.GetTestPlan(ctx)
}
