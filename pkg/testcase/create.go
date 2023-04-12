package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	relatedtestcase "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/relatedtestcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"

	commonpb "github.com/NpoolPlatform/message/npool"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
)

func CreateTestCase(ctx context.Context, handler *Handler) (*npool.TestCase, error) {
	info, err := testcasemwcli.CreateTestCase(
		ctx,
		&testcasemwpb.CreateTestCaseReq{
			Name:              *handler.Name,
			Description:       handler.Description,
			ModuleName:        handler.ModuleName,
			ApiID:             *handler.ApiID,
			Arguments:         *handler.Arguments,
			ExpectationResult: *handler.ExpectationResult,
		},
	)
	if err != nil {
		logger.Sugar().Errorw("CreateTestCase", "err", err)
		return nil, err
	}

	_api, err := apimwcli.GetAPIOnly(
		ctx,
		&mgrpb.Conds{
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
		ID:                info.ID,
		Name:              info.Name,
		ModuleID:          info.ModuleID,
		ModuleName:        info.ModuleName,
		ApiID:             info.ApiID,
		ApiPath:           _api.Path,
		ApiPathPrefix:     _api.PathPrefix,
		ApiServiceName:    _api.ServiceName,
		ApiProtocol:       _api.Protocol.String(),
		ApiMethod:         _api.Method.String(),
		ApiDeprecated:     _api.GetDepracated(),
		ApiCreatedAt:      _api.CreatedAt,
		ApiUpdatedAt:      _api.UpdatedAt,
		Arguments:         info.Arguments,
		ExpectationResult: info.ExpectationResult,
		TestCaseType:      info.TestCaseType,
		RelatedTestCases:  []*relatedtestcase.RelatedTestCase{},
		Deprecated:        info.Deprecated,
		CreatedAt:         info.CreatedAt,
		UpdatedAt:         info.UpdatedAt,
	}

	return _info, nil
}
