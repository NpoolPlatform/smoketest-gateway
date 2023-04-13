package testplan

import (
	"context"
	"fmt"

	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID                *string
	Name              *string
	State             *string
	OwnerID           *string
	ResponsibleUserID *string
	Deadline          *uint32
	Conds             *testplanmgrpb.Conds
	Offset            *int32
	Limit             *int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.ID = id
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		h.Name = name
		return nil
	}
}

func WithOwnerID(ownerID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*ownerID); err != nil {
			return err
		}
		h.OwnerID = ownerID
		return nil
	}
}

func WithDeadline(deadline *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deadline == nil {
			return nil
		}
		h.Deadline = deadline
		return nil
	}
}

func WithResponsibleUserID(userID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			return nil
		}
		if _, err := uuid.Parse(*userID); err != nil {
			return err
		}
		h.ResponsibleUserID = userID
		return nil
	}
}

func WithConds(conds *testplanmgrpb.Conds, offset, limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return fmt.Errorf("invalid conds")
		}

		if conds.ID != nil {
			if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
				return err
			}
		}
		if conds.ResponsibleUserID != nil {
			if _, err := uuid.Parse(conds.GetResponsibleUserID().GetValue()); err != nil {
				return err
			}
		}

		h.Conds = conds

		if h.Offset == nil {
			offset = constant.DefaultRowLimit
		}
		h.Offset = &offset
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = &limit

		return nil
	}
}
