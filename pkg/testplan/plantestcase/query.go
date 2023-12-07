package plantestcase

import (
	"context"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	apimwpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan/plantestcase"
	plantestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan/plantestcase"
)

type queryHandler struct {
	*Handler
	testCases []*plantestcasemwpb.PlanTestCase
	apis      map[string]*apimwpb.API
	users     map[string]*usermwpb.User
	infos     []*npool.PlanTestCase
}

func (h *queryHandler) getAPIs(ctx context.Context) error {
	apiIDs := []string{}
	for _, info := range h.testCases {
		apiIDs = append(apiIDs, info.TestCaseApiID)
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

func (h *queryHandler) getUsers(ctx context.Context) error {
	userIDs := []string{}
	for _, info := range h.testCases {
		userIDs = append(userIDs, info.TestUserID)
	}
	users, _, err := usermwcli.GetUsers(ctx, &usermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}, 0, int32(len(userIDs)))
	if err != nil {
		return err
	}
	for _, user := range users {
		h.users[user.EntID] = user
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.testCases {
		_info := &npool.PlanTestCase{
			ID:            info.ID,
			EntID:         info.EntID,
			TestPlanID:    info.TestPlanID,
			TestCaseID:    info.TestCaseID,
			TestCaseName:  info.TestCaseName,
			TestCaseType:  info.TestCaseType,
			TestCaseClass: info.TestCaseClass,
			ApiID:         info.TestCaseApiID,
			Input:         info.Input,
			Output:        info.Output,
			Description:   info.Description,
			RunDuration:   info.RunDuration,
			TestUserID:    info.TestUserID,
			Result:        info.Result,
			Index:         info.Index,
			ModuleID:      info.ModuleID,
			ModuleName:    info.ModuleName,
			CreatedAt:     info.CreatedAt,
			UpdatedAt:     info.UpdatedAt,
		}
		if api, ok := h.apis[info.TestCaseApiID]; ok {
			_info.ApiPath = api.Path
		}
		if user, ok := h.users[info.TestUserID]; ok {
			_info.TestUserEmail = user.EmailAddress
		}

		h.infos = append(h.infos, _info)
	}
}

func (h *Handler) GetPlanTestCases(ctx context.Context) ([]*npool.PlanTestCase, uint32, error) {
	infos, total, err := cli.GetPlanTestCases(ctx, &plantestcasemwpb.Conds{
		TestPlanID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.TestPlanID},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:   h,
		testCases: infos,
		apis:      map[string]*apimwpb.API{},
	}
	if err := handler.getAPIs(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, total, nil
}

func (h *Handler) GetPlanTestCase(ctx context.Context) (*npool.PlanTestCase, error) {
	info, err := cli.GetPlanTestCase(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}

	handler := &queryHandler{
		Handler:   h,
		testCases: []*plantestcasemwpb.PlanTestCase{info},
		apis:      map[string]*apimwpb.API{},
	}
	if err := handler.getAPIs(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetPlanTestCaseExt(ctx context.Context, info *plantestcasemwpb.PlanTestCase) (*npool.PlanTestCase, error) {
	handler := &queryHandler{
		Handler:   h,
		testCases: []*plantestcasemwpb.PlanTestCase{info},
		apis:      map[string]*apimwpb.API{},
	}
	if err := handler.getAPIs(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
