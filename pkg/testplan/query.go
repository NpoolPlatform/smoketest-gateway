package testplan

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	testplanmwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

func (handler *Handler) GetTestPlans(ctx context.Context) ([]*npool.TestPlan, uint32, error) {
	infos, total, err := testplanmwcli.GetTestPlans(
		ctx,
		&testplanmgrpb.Conds{},
		*handler.Offset,
		*handler.Limit,
	)
	if err != nil {
		logger.Sugar().Errorw("GetTestPlans", "err", err)
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.TestPlan{}, 0, nil
	}

	_infos := []*npool.TestPlan{}
	for _, info := range infos {
		row := npool.TestPlan{
			ID:                    info.ID,
			Name:                  info.Name,
			State:                 testplanmgrpb.TestPlanState(info.GetState()),
			OwnerID:               info.OwnerID,
			OwnerName:             info.OwnerID,
			ResponsibleUserID:     info.ResponsibleUserID,
			ResponsibleUsername:   info.ResponsibleUserID,
			FailedTestCasesCount:  info.FailedTestCasesCount,
			SkippedTestCasesCount: info.SkippedTestCasesCount,
			PassedTestCasesCount:  info.PassedTestCasesCount,
			RunDuration:           info.RunDuration,
			TestResult:            testplanmgrpb.TestResultState(info.TestResult),
			Deadline:              info.Deadline,
			CreatedAt:             info.CreatedAt,
		}
		_infos = append(_infos, &row)
	}

	return _infos, total, nil
}

func (handler *Handler) GetTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	if handler.ID == nil {
		return nil, fmt.Errorf("invalid testplan id")
	}

	info, err := testplanmwcli.GetTestPlan(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}

	_info := &npool.TestPlan{
		ID:                    info.ID,
		Name:                  info.Name,
		State:                 testplanmgrpb.TestPlanState(info.GetState()),
		OwnerID:               info.OwnerID,
		OwnerName:             info.OwnerID,
		ResponsibleUserID:     info.ResponsibleUserID,
		ResponsibleUsername:   info.ResponsibleUserID,
		FailedTestCasesCount:  info.FailedTestCasesCount,
		SkippedTestCasesCount: info.SkippedTestCasesCount,
		PassedTestCasesCount:  info.PassedTestCasesCount,
		RunDuration:           info.RunDuration,
		TestResult:            testplanmgrpb.TestResultState(info.TestResult),
		Deadline:              info.Deadline,
		CreatedAt:             info.CreatedAt,
	}
	return _info, nil
}
