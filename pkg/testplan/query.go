package testplan

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	ptcpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
	ptccli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

type queryHandler struct {
	*Handler
	infos []*pb.TestPlan
}

//nolint
func (h *queryHandler) formalize(ctx context.Context) ([]*npool.TestPlan, error) {
	planIDs := []string{}
	// userIDs := []string{}
	for _, info := range h.infos {
		planIDs = append(planIDs, info.ID)
		// userIDs = append(userIDs, info.Executor, info.CreatedBy)
	}

	_infos, total, err := ptccli.GetPlanTestCases(ctx, &ptcpb.Conds{
		TestPlanIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: planIDs,
		},
	}, 0, 1)
	if err != nil {
		return nil, err
	}

	if len(_infos) < int(total) {
		_infos, _, err = ptccli.GetPlanTestCases(ctx, &ptcpb.Conds{
			TestPlanIDs: &commonpb.StringSliceVal{
				Op:    cruder.IN,
				Value: planIDs,
			},
		}, 0, int32(total))
		if err != nil {
			return nil, err
		}
	}

	planTestCaseMap := map[string][]*ptcpb.PlanTestCase{}
	skipMap := map[string]int{}
	passMap := map[string]int{}
	failMap := map[string]int{}
	for _, info := range _infos {
		_, ok := planTestCaseMap[info.TestPlanID]
		if !ok {
			planTestCaseMap[info.TestPlanID] = []*ptcpb.PlanTestCase{}
		}
		rows := planTestCaseMap[info.TestPlanID]
		rows = append(rows, info)
		planTestCaseMap[info.TestPlanID] = rows

		_, ok = skipMap[info.TestPlanID]
		if !ok {
			skipMap[info.TestPlanID] = 0
		}
		_, ok = passMap[info.TestPlanID]
		if !ok {
			passMap[info.TestPlanID] = 0
		}
		_, ok = failMap[info.TestPlanID]
		if !ok {
			failMap[info.TestPlanID] = 0
		}

		switch info.Result {
		case ptcpb.TestCaseResult_Skipped:
			skipMap[info.TestPlanID] += 1
		case ptcpb.TestCaseResult_Passed:
			passMap[info.TestPlanID] += 1
		case ptcpb.TestCaseResult_Failed:
			failMap[info.TestPlanID] += 1
		default:
			//pass
			continue
		}
	}

	infos := []*npool.TestPlan{}
	for _, info := range h.infos {
		planTestCases, ok := planTestCaseMap[info.ID]
		if !ok {
			continue
		}
		skips, ok := skipMap[info.ID]
		if !ok {
			continue
		}
		passes, ok := passMap[info.ID]
		if !ok {
			continue
		}
		fails, ok := failMap[info.ID]
		if !ok {
			continue
		}
		row := npool.TestPlan{
			ID:               info.ID,
			Name:             info.Name,
			State:            info.GetState(),
			CreatedBy:        info.CreatedBy,
			Username:         info.CreatedBy,
			Executor:         info.Executor,
			ExecutorUsername: info.Executor,
			Fails:            uint32(fails),
			Skips:            uint32(skips),
			Passes:           uint32(passes),
			RunDuration:      info.RunDuration,
			Result:           info.Result,
			Deadline:         info.Deadline,
			CreatedAt:        info.CreatedAt,
			PlanTestCases:    planTestCases,
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
