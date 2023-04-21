package testplan

import (
	"context"
	"fmt"

	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	constant "github.com/NpoolPlatform/smoketest-gateway/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID        *string
	Name      *string
	State     **testplanmgrpb.TestPlanState
	CreatedBy *string
	Executor  *string
	Deadline  *uint32
	Offset    int32
	Limit     int32
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
		const leastNameLen = 4
		if len(*name) < leastNameLen {
			return fmt.Errorf("name %v too short", *name)
		}
		h.Name = name
		return nil
	}
}

func WithCreatedBy(createdBy *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*createdBy); err != nil {
			return err
		}
		h.CreatedBy = createdBy
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

func WithExecutor(executor *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if executor == nil {
			return nil
		}
		if _, err := uuid.Parse(*executor); err != nil {
			return err
		}
		h.Executor = executor
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
