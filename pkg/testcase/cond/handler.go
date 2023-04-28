package cond

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *string
	TestCaseID     *string
	CondTestCaseID *string
	CondType       *pb.CondType
	Index          *uint32
	ArgumentMap    *string
	Offset         int32
	Limit          int32
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

func WithTestCaseID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}

		if _, err := cli.ExistTestCase(ctx, *id); err != nil {
			return err
		}

		h.TestCaseID = id
		return nil
	}
}

func WithCondTestCaseID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		if _, err := cli.ExistTestCase(ctx, *id); err != nil {
			return err
		}

		h.CondTestCaseID = id
		return nil
	}
}

func WithCondType(_type *pb.CondType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case pb.CondType_PreCondition:
		case pb.CondType_Cleaner:
		default:
			return fmt.Errorf("invalid CondType")
		}

		h.CondType = _type
		return nil
	}
}

func WithIndex(index *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if index == nil {
			return nil
		}
		h.Index = index
		return nil
	}
}

func WithArgumentMap(argMap *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if argMap == nil {
			return nil
		}

		var r interface{}
		if err := json.Unmarshal([]byte(*argMap), &r); err != nil {
			return err
		}
		h.ArgumentMap = argMap
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
