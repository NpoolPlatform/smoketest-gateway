package plantestcase

import (
	"context"
	"fmt"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	plantestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
	plantestcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

func (handler *Handler) Getplantestcases(ctx context.Context) ([]*npool.plantestcase, uint32, error) {
	infos, total, err := plantestcasemwcli.Getplantestcases(
		ctx,
		&plantestcasemwpb.Conds{
			TestPlanID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: *handler.TestPlanID,
			},
		},
		*handler.Offset,
		*handler.Limit,
	)
	if err != nil {
		logger.Sugar().Errorw("Getplantestcases", "err", err)
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.plantestcase{}, 0, nil
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

	_infos := []*npool.plantestcase{}
	for _, info := range infos {
		_testCase, ok := testCaseMap[info.TestCaseID]
		if !ok {
			continue
		}

		_api, ok := apiMap[_testCase.ApiID]
		if !ok {
			continue
		}

		row := npool.plantestcase{
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
			TestCaseResult:      plantestcasemwpb.TestCaseResult_DefaultTestCaseResult,
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

func (handler *Handler) Getplantestcase(ctx context.Context) (*npool.plantestcase, error) {
	if handler.ID == nil {
		return nil, fmt.Errorf("invalid plantestcase id")
	}

	info, err := plantestcasemwcli.Getplantestcase(ctx, *handler.ID)
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

	_info := &npool.plantestcase{
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
		TestCaseResult:      plantestcasemwpb.TestCaseResult_DefaultTestCaseResult,
		TestCaseExpectation: _testCase.ExpectationResult,
		Description:         info.Description,
		RunDuration:         info.RunDuration,
		TestUserID:          info.TestCaseID,
		TestUsername:        info.TestUserID,
		CreatedAt:           info.CreatedAt,
	}
	return _info, nil
}
