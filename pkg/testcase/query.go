package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	apipb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

type queryHandler struct {
	*Handler
	Infos []*pb.TestCase
}

func (h *queryHandler) formalize(ctx context.Context) ([]*npool.TestCase, error) {
	apiIDs := []string{}
	for _, info := range h.Infos {
		apiIDs = append(apiIDs, info.ApiID)
	}

	apis, _, err := apicli.GetAPIs(ctx, &apipb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: apiIDs},
	}, int32(len(apiIDs)), 0)
	if err != nil {
		return nil, err
	}

	apiMap := map[string]*apipb.API{}
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
			Description:    info.Description,
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
			Expectation:    info.Expectation,
			OutputDesc:     info.OutputDesc,
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

	conds := &pb.Conds{}
	if handler.ModuleID != nil {
		conds.ModuleID = &basetypes.StringVal{Op: cruder.EQ, Value: *handler.ModuleID}
	}

	infos, total, err := cli.GetTestCases(ctx, conds, handler.Offset, handler.Limit)
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

	info, err := cli.GetTestCase(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	handler := &queryHandler{
		Handler: h,
	}

	handler.Infos = []*pb.TestCase{info}

	_infos, err := handler.formalize(ctx)
	if err != nil {
		return nil, err
	}

	return _infos[0], nil
}
