package testplan

import (
	"context"
	"fmt"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testplan"
	testplanmwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplanmwcli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
)

type queryHandler struct {
	*Handler
	testPlans []*testplanmwpb.TestPlan
	infos     []*npool.TestPlan
	users     map[string]*usermwpb.User
}

func (h *queryHandler) getUsers(ctx context.Context) error {
	userIDs := []string{}
	for _, info := range h.testPlans {
		userIDs = append(userIDs, info.CreatedBy, info.Executor)
	}
	users, _, err := usermwcli.GetUsers(ctx, &usermwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}, 0, int32(len(userIDs)))
	if err != nil {
		return err
	}
	for _, user := range users {
		h.users[user.ID] = user
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.testPlans {
		creator, ok := h.users[info.CreatedBy]
		if !ok {
			continue
		}
		executor, ok := h.users[info.Executor]
		if !ok {
			continue
		}
		row := npool.TestPlan{
			ID:            info.ID,
			Name:          info.Name,
			State:         info.State,
			CreatedBy:     info.CreatedBy,
			CreatorEmail:  creator.EmailAddress,
			Executor:      info.Executor,
			ExecutorEmail: executor.EmailAddress,
			Fails:         info.Fails,
			Skips:         info.Skips,
			Passes:        info.Passes,
			RunDuration:   info.RunDuration,
			Result:        info.Result,
			Deadline:      info.Deadline,
			CreatedAt:     info.CreatedAt,
		}
		h.infos = append(h.infos, &row)
	}
}

func (h *Handler) GetTestPlans(ctx context.Context) ([]*npool.TestPlan, uint32, error) {
	infos, total, err := testplanmwcli.GetTestPlans(ctx, nil, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	handler := &queryHandler{
		Handler:   h,
		testPlans: infos,
		users:     map[string]*usermwpb.User{},
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, total, nil
}

func (h *Handler) GetTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := testplanmwcli.GetTestPlan(ctx, *h.ID)
	if err != nil {
		return nil, err
	}

	handler := &queryHandler{
		Handler:   h,
		testPlans: []*testplanmwpb.TestPlan{info},
		users:     map[string]*usermwpb.User{},
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
