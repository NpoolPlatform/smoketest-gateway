package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase/cond"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"

	commonpb "github.com/NpoolPlatform/message/npool"
)

type queryHandler struct {
	*Handler
	moduleID string
}

func (h *queryHandler) validate(multiple bool) error {
	if multiple {
		if h.Offset == nil {
			return fmt.Errorf("invalid offset")
		}
		if h.Limit == nil {
			return fmt.Errorf("invalid limit")
		}
	} else {
		if h.ID == nil {
			return fmt.Errorf("invalid id")
		}
	}

	return nil
}

func (h *queryHandler) formalize(infos []*testcasemgrpb.TestCase) ([]*npool.TestCase, error) {
	apiIDs := []string{}
	for _, info := range infos {
		apiIDs = append(apiIDs, info.ApiID)
	}

	apis, _, err := apimwcli.GetAPIs(ctx, &apimgrpb.Conds{
		IDs: &commonpb.StringSliceVal{Op: cruder.IN, Value: apiIDs},
	}, int32(len(apiIDs)), 0)
	if err != nil {
		return nil, err
	}

	apiMap := map[string]*apimgrpb.API{}
	for _, row := range apis {
		apiMap[row.ID] = row
	}

	_infos := []*npool.TestCase{}
	for _, info := range infos {
		_api, ok := apiMap[info.ApiID]
		if !ok {
			continue
		}
		row := npool.TestCase{
			ID:                 info.ID,
			Name:               info.Name,
			ModuleID:           info.ModuleID,
			ModuleName:         info.ModuleName,
			ApiID:              info.ApiID,
			ApiPath:            _api.Path,
			ApiPathPrefix:      _api.PathPrefix,
			ApiServiceName:     _api.ServiceName,
			ApiProtocol:        _api.Protocol.String(),
			ApiMethod:          _api.Method.String(),
			ApiDeprecated:      _api.Depracated,
			Arguments:          info.Arguments,
			ArgTypeDescription: info.ArgTypeDescription,
			TestCaseType:       info.TestCaseType,
			RelatedTestCases:   []*cond.RelatedTestCase{},
			Deprecated:         info.Deprecated,
			CreatedAt:          info.CreatedAt,
			UpdatedAt:          info.UpdatedAt,
		}
		_infos = append(_infos, &row)
	}

	return _infos, nil
}

func (h *Handler) GetTestCases(ctx context.Context) ([]*npool.TestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	if err := handler.validate(true); err != nil {
		return nil, 0, err
	}

	conds := &testcasemgrpb.Conds{}
	if h.ModuleID != nil {
		conds.ModuleID = &commonpb.StringVal{Op: cruder.EQ, Value: *h.ModuleID}
	}

	infos, total, err := testcasemwcli.GetTestCases(ctx, conds, *handler.Offset, *handler.Limit)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return nil, 0, nil
	}

	_infos, err := h.formalize(infos)
	if err != nil {
		return nil, 0, err
	}

	return _infos, total, nil
}

func (handler *Handler) GetTestCase(ctx context.Context) (*npool.TestCase, error) {
	if err := handler.validate(true); err != nil {
		return nil, err
	}

	info, err := testcasemwcli.GetTestCase(ctx, *handler.ID)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid testcase id")
	}

	_infos, err := h.formalize([]*testcasemgrpb.TestCase{info})
	if err != nil {
		return nil, err
	}

	return _infos[0], nil
}
