package planrelatedtestcase

import (
	"context"
	"fmt"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/planrelatedtestcase"
	"github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/planrelatedtestcase"
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	planrelatedtestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	planrelatedtestcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/planrelatedtestcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

func (handler *Handler) GetPlanRelatedTestCases(ctx context.Context) ([]*npool.PlanRelatedTestCase, uint32, error) {
	infos, total, err := planrelatedtestcasemwcli.GetPlanRelatedTestCases(
		ctx,
		&planrelatedtestcasemwpb.Conds{
			TestPlanID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: *handler.TestPlanID,
			},
		},
		*handler.Offset,
		*handler.Limit,
	)
	if err != nil {
		logger.Sugar().Errorw("GetPlanRelatedTestCases", "err", err)
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.PlanRelatedTestCase{}, 0, nil
	}

	testCaseIDs := []string{}
	for _, info := range infos {
		testCaseIDs = append(testCaseIDs, info.TestCaseID)
	}

	testCases, _, err := testcasemwcli.GetTestCases(ctx, &mgrpb.Conds{
		IDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: testCaseIDs,
		},
	}, 0, int32(len(testCaseIDs)))
	if err != nil {
		return nil, 0, err
	}

	testCaseMap := map[string]*testcasemwpb.TestCase{}

	apiIDs := []string{}
	for _, row := range testCases {
		testCaseMap[row.ID] = row
		apiIDs = append(apiIDs, row.ApiID)
	}

	apis, _, err := apimwcli.GetAPIs(ctx, &apimgrpb.Conds{
		IDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: apiIDs,
		},
	}, int32(len(apiIDs)), 0)
	if err != nil {
		logger.Sugar().Errorw("GetTestCases", "err", err)
		return nil, 0, err
	}

	apiMap := map[string]*apimgrpb.API{}

	for _, row := range apis {
		apiMap[row.ID] = row
	}

	_infos := []*npool.PlanRelatedTestCase{}
	for _, info := range infos {
		_testCase, ok := testCaseMap[info.TestCaseID]
		if !ok {
			continue
		}

		_api, ok := apiMap[_testCase.ApiID]
		if !ok {
			continue
		}

		row := npool.PlanRelatedTestCase{
			ID:                  info.ID,
			TestPlanID:          info.TestPlanID,
			TestPlanName:        info.TestPlanID,
			TestCaseID:          info.TestCaseID,
			ApiID:               _testCase.ApiID,
			ApiPath:             _api.Path,
			ApiPathPrefix:       _api.PathPrefix,
			TestCaseArguments:   _testCase.Arguments,
			TestCaseType:        _testCase.TestCaseType,
			TestCaseOutput:      info.TestCaseOutput,
			TestCaseResult:      planrelatedtestcase.TestCaseResult_DefaultTestCaseResult,
			TestCaseExpectation: _testCase.ExpectationResult,
			Description:         info.Description,
			RunDuration:         info.RunDuration,
			TestUserID:          info.TestCaseID,
			TestUsername:        info.TestUserID,
			CreatedAt:           info.CreatedAt,
		}
		_infos = append(_infos, &row)
	}

	return _infos, total, nil
}

func (handler *Handler) GetPlanRelatedTestCase(ctx context.Context) (*npool.PlanRelatedTestCase, error) {
	if handler.ID == nil {
		return nil, fmt.Errorf("invalid planrelatedtestcase id")
	}

	info, err := planrelatedtestcasemwcli.GetPlanRelatedTestCase(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}

	_testCase, err := testcasemwcli.GetTestCase(ctx, info.TestCaseID)
	if err != nil {
		return nil, err
	}

	_api, err := apimwcli.GetAPIOnly(ctx, &apimgrpb.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: _testCase.ApiID,
		},
	})
	if err != nil {
		return nil, err
	}

	_info := &npool.PlanRelatedTestCase{
		ID:                  info.ID,
		TestPlanID:          info.TestPlanID,
		TestPlanName:        info.TestPlanID,
		TestCaseID:          info.TestCaseID,
		ApiID:               _testCase.ApiID,
		ApiPath:             _api.Path,
		ApiPathPrefix:       _api.PathPrefix,
		TestCaseArguments:   _testCase.Arguments,
		TestCaseType:        _testCase.TestCaseType,
		TestCaseOutput:      info.TestCaseOutput,
		TestCaseResult:      planrelatedtestcase.TestCaseResult_DefaultTestCaseResult,
		TestCaseExpectation: _testCase.ExpectationResult,
		Description:         info.Description,
		RunDuration:         info.RunDuration,
		TestUserID:          info.TestCaseID,
		TestUsername:        info.TestUserID,
		CreatedAt:           info.CreatedAt,
	}
	return _info, nil
}
