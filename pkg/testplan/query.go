package testplan

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

type queryHandler struct {
	*Handler
	infos []*pb.TestPlan
}

//nolint
func (h *queryHandler) formalize(ctx context.Context) ([]*npool.TestPlan, error) {
	// userIDs := []string{}
	// for _, info := range h.infos {
	// userIDs = append(userIDs, info.Executor, info.CreatedBy)
	// }

	// appusercli.GetUserOnly(ctx, &appuserpb.C)

	infos := []*npool.TestPlan{}
	for _, info := range h.infos {
		// planTestCases, ok := planTestCaseMap[info.ID]
		// if !ok {
		// 	continue
		// }
		row := npool.TestPlan{
			ID:            info.ID,
			Name:          info.Name,
			State:         info.GetState(),
			CreatedBy:     info.CreatedBy,
			Email:         info.CreatedBy,
			Executor:      info.Executor,
			ExecutorEmail: info.Executor,
			Fails:         info.Fails,
			Skips:         info.Skips,
			Passes:        info.Passes,
			RunDuration:   info.RunDuration,
			Result:        info.Result,
			Deadline:      info.Deadline,
			CreatedAt:     info.CreatedAt,
		}
		infos = append(infos, &row)
	}

	return infos, nil
}

func (h *Handler) GetTestPlans(ctx context.Context) ([]*npool.TestPlan, uint32, error) {
	infos, total, err := cli.GetTestPlans(ctx, nil, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.TestPlan{}, 0, nil
	}

	handler := &queryHandler{
		Handler: h,
	}

	handler.infos = infos
	_infos, err := handler.formalize(ctx)
	if err != nil {
		return nil, 0, err
	}

	return _infos, total, nil
}

func (h *Handler) GetTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := cli.GetTestPlan(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	handler := &queryHandler{
		Handler: h,
	}

	handler.infos = []*pb.TestPlan{info}

	_info, err := handler.formalize(ctx)
	if err != nil {
		return nil, err
	}

	return _info[0], nil
}
