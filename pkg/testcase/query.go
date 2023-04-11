package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/relatedtestcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

func GetTestCases(ctx context.Context, offset, limit int32) ([]*npool.TestCase, uint32, error) {
	infos, total, err := testcasemwcli.GetTestCases(
		ctx,
		&testcasemgrpb.Conds{},
		offset,
		limit,
	)
	if err != nil {
		logger.Sugar().Errorw("GetTestCases", "err", err)
		return nil, 0, err
	}

	if len(infos) == 0 {
		return []*npool.TestCase{}, 0, nil
	}

	testCaseIDs := []string{}
	apiIDs := []string{}
	for _, info := range infos {
		apiIDs = append(apiIDs, info.ApiID)
		testCaseIDs = append(testCaseIDs, info.ID)
	}

	fmt.Println("TestCaseIDs: ", testCaseIDs)
	fmt.Println("ApiIDs: ", apiIDs)

	// 查询用例关联的前置条件|后置条件

	// TODO

	// 查询API信息

	// TODO

	_infos := []*npool.TestCase{}
	_relatedTestCases := []*relatedtestcase.RelatedTestCase{}
	for _, info := range infos {
		row := npool.TestCase{
			ID:                info.ID,
			Name:              info.Name,
			ModuleID:          info.ModuleID,
			ModuleName:        info.ModuleName,
			ApiID:             info.ApiID,
			ApiPath:           "",
			ApiPathPrefix:     "",
			ApiServiceName:    "",
			ApiProtocol:       "",
			ApiMethod:         "",
			ApiDeprecated:     "",
			ApiCreatedAt:      0,
			ApiUpdatedAt:      0,
			Arguments:         info.Arguments,
			ExpectationResult: info.ExpectationResult,
			TestCaseType:      info.TestCaseType,
			RelatedTestCases:  _relatedTestCases,
			Deprecated:        info.Deprecated,
			CreatedAt:         info.CreatedAt,
			UpdatedAt:         info.UpdatedAt,
		}
		_infos = append(_infos, &row)
	}

	return _infos, total, nil

}
