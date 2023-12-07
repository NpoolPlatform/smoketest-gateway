package testcase

import (
	"context"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	apimwpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

type queryHandler struct {
	*Handler
	testCases []*testcasemwpb.TestCase
	infos     []*npool.TestCase
	apis      map[string]*apimwpb.API
}

func (h *queryHandler) getAPIs(ctx context.Context) error {
	apiIDs := []string{}
	for _, info := range h.testCases {
		apiIDs = append(apiIDs, info.ApiID)
	}
	apis, _, err := apimwcli.GetAPIs(ctx, &apimwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: apiIDs},
	}, 0, int32(len(apiIDs)))
	if err != nil {
		return err
	}
	for _, api := range apis {
		h.apis[api.EntID] = api
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.testCases {
		_info := npool.TestCase{
			ID:            info.ID,
			EntID:         info.EntID,
			Name:          info.Name,
			Description:   info.Description,
			ModuleID:      info.ModuleID,
			ModuleName:    info.ModuleName,
			ApiID:         info.ApiID,
			Input:         info.Input,
			InputDesc:     info.InputDesc,
			Expectation:   info.Expectation,
			OutputDesc:    info.OutputDesc,
			TestCaseType:  info.TestCaseType,
			TestCaseClass: info.TestCaseClass,
			Deprecated:    info.Deprecated,
			CreatedAt:     info.CreatedAt,
			UpdatedAt:     info.UpdatedAt,
		}
		api, ok := h.apis[info.ApiID]
		if ok {
			_info.ApiPath = api.Path
			_info.ApiPathPrefix = api.PathPrefix
			_info.ApiServiceName = api.ServiceName
			_info.ApiProtocol = api.Protocol.String()
			_info.ApiMethod = api.Method.String()
			_info.ApiDeprecated = api.Deprecated
		}
		h.infos = append(h.infos, &_info)
	}
}

func (h *Handler) GetTestCases(ctx context.Context) ([]*npool.TestCase, uint32, error) {
	conds := &testcasemwpb.Conds{}
	if h.ModuleID != nil {
		conds.ModuleID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.ModuleID}
	}
	infos, total, err := testcasemwcli.GetTestCases(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	handler := &queryHandler{
		Handler:   h,
		testCases: infos,
		apis:      map[string]*apimwpb.API{},
	}
	if err := handler.getAPIs(ctx); err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, total, nil
}

func (h *Handler) GetTestCase(ctx context.Context) (*npool.TestCase, error) {
	info, err := testcasemwcli.GetTestCase(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}

	handler := &queryHandler{
		Handler:   h,
		testCases: []*testcasemwpb.TestCase{info},
		apis:      map[string]*apimwpb.API{},
	}
	if err := handler.getAPIs(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetTestCaseExt(ctx context.Context, info *testcasemwpb.TestCase) (*npool.TestCase, error) {
	handler := &queryHandler{
		Handler:   h,
		testCases: []*testcasemwpb.TestCase{info},
		apis:      map[string]*apimwpb.API{},
	}
	if err := handler.getAPIs(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
