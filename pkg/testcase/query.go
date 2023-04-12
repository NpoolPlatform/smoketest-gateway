package testcase

import (
	"context"
	"fmt"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
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

	// TODO:After

	// 查询API信息

	// TODO:After

	_infos := []*npool.TestCase{}
	_relatedTestCases := []*relatedtestcase.RelatedTestCase{}
	for _, info := range infos {
		row := npool.TestCase{
			ID:               info.ID,
			Name:             info.Name,
			ModuleID:         info.ModuleID,
			ModuleName:       info.ModuleName,
			ApiID:            info.ApiID,
			ApiPath:          "",
			ApiPathPrefix:    "",
			ApiServiceName:   "",
			ApiProtocol:      "",
			ApiMethod:        "",
			ApiDeprecated:    false,
			Arguments:        info.Arguments,
			TestCaseType:     info.TestCaseType,
			RelatedTestCases: _relatedTestCases,
			Deprecated:       info.Deprecated,
			CreatedAt:        info.CreatedAt,
			UpdatedAt:        info.UpdatedAt,
		}
		_infos = append(_infos, &row)
	}

	return _infos, total, nil
}

func (handler *Handler) GetTestCase(ctx context.Context) (*npool.TestCase, error) {
	if handler.ID == nil {
		return nil, fmt.Errorf("invalid testcase id")
	}

	info, err := testcasemwcli.GetTestCase(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}

	_api, err := apimwcli.GetAPIOnly(
		ctx,
		&apimgrpb.Conds{
			ID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: info.ApiID,
			},
		},
	)
	if err != nil {
		logger.Sugar().Errorw("CreateTestCase", "err", err)
		return nil, err
	}

	_info := &npool.TestCase{
		ID:               info.ID,
		Name:             info.Name,
		ModuleID:         info.ModuleID,
		ModuleName:       info.ModuleName,
		ApiID:            info.ApiID,
		ApiPath:          _api.Path,
		ApiPathPrefix:    _api.PathPrefix,
		ApiServiceName:   _api.ServiceName,
		ApiProtocol:      _api.Protocol.String(),
		ApiMethod:        _api.Method.String(),
		ApiDeprecated:    _api.GetDepracated(),
		Arguments:        info.Arguments,
		Expectation:      info.ExpectationResult,
		TestCaseType:     info.TestCaseType,
		RelatedTestCases: []*relatedtestcase.RelatedTestCase{},
		Deprecated:       info.Deprecated,
		CreatedAt:        info.CreatedAt,
		UpdatedAt:        info.UpdatedAt,
	}
	return _info, nil
}
