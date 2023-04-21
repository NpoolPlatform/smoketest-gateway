package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasemwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"

	commonpb "github.com/NpoolPlatform/message/npool"
)

type queryHandler struct {
	*Handler
	Infos []*testcasemwpb.TestCase
}

func (h *queryHandler) formalize(ctx context.Context) ([]*npool.TestCase, error) {
	apiIDs := []string{}
	for _, info := range h.Infos {
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

	infos := []*npool.TestCase{}
	for _, info := range h.Infos {
		api, ok := apiMap[info.ApiID]
		if !ok {
			continue
		}
		row := npool.TestCase{
			ID:             info.ID,
			Name:           info.Name,
			ModuleID:       info.ModuleID,
			ModuleName:     info.ModuleName,
			ApiID:          info.ApiID,
			ApiPath:        api.Path,
			ApiPathPrefix:  api.PathPrefix,
			ApiServiceName: api.ServiceName,
			ApiProtocol:    api.Protocol.String(),
			ApiMethod:      api.Method.String(),
			ApiDeprecated:  api.Depracated,
			Input:          info.Input,
			InputDesc:      info.InputDesc,
			TestCaseType:   info.TestCaseType,
			Deprecated:     info.Deprecated,
			CreatedAt:      info.CreatedAt,
			UpdatedAt:      info.UpdatedAt,
		}
		infos = append(infos, &row)
	}

	return infos, nil
}

func (h *Handler) GetTestCases(ctx context.Context) ([]*npool.TestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	conds := &testcasemgrpb.Conds{}
	if handler.ModuleID != nil {
		conds.ModuleID = &commonpb.StringVal{Op: cruder.EQ, Value: *handler.ModuleID}
	}

	infos, total, err := testcasemwcli.GetTestCases(ctx, conds, handler.Offset, handler.Limit)
	if err != nil {
		return nil, 0, err
	}

	if len(infos) == 0 {
		return nil, 0, nil
	}

	handler.Infos = infos

	_infos, err := handler.formalize(ctx)
	if err != nil {
		return nil, 0, err
	}

	return _infos, total, nil
}

func (h *Handler) GetTestCase(ctx context.Context) (*npool.TestCase, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := testcasemwcli.GetTestCase(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, fmt.Errorf("invalid testcase")
	}

	handler := &queryHandler{
		Handler: h,
	}

	handler.Infos = []*testcasemwpb.TestCase{info}

	_infos, err := handler.formalize(ctx)
	if err != nil {
		return nil, err
	}

	return _infos[0], nil
}
