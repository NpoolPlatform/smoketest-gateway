package plantestcase

import (
	"context"
	"fmt"

	plantestcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *string
	TestPlanID     *string
	TestCaseID     *string
	TestUserID     *string
	TestCaseOutput *string
	TestCaseResult *plantestcasemwpb.TestCaseResult
	Description    *string
	Index          *uint32
	RunDuration    *uint32
	Conds          *plantestcasemwpb.Conds
	Offset         *int32
	Limit          *int32
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

func WithTestPlanID(planID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*planID); err != nil {
			return err
		}
		h.TestPlanID = planID
		return nil
	}
}

func WithTestCaseID(testCaseID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*testCaseID); err != nil {
			return err
		}
		h.TestCaseID = testCaseID
		return nil
	}
}

func WithTestUserID(userID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			return nil
		}
		if _, err := uuid.Parse(*userID); err != nil {
			return err
		}
		h.TestUserID = userID
		return nil
	}
}

func WithTestCaseOutput(output *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if output == nil {
			return nil
		}
		h.TestCaseOutput = output
		return nil
	}
}

func WithTestCaseResult(result *plantestcasemwpb.TestCaseResult) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			return fmt.Errorf("need testcase result")
		}
		h.TestCaseResult = result
		return nil
	}
}

func WithRunDuration(duration *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			return nil
		}
		h.RunDuration = duration
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

func WithDescription(description *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			return nil
		}
		h.Description = description
		return nil
	}
}

func WithConds(conds *plantestcasemwpb.Conds, offset, limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return fmt.Errorf("invalid conds")
		}

		if conds.ID != nil {
			if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
				return err
			}
		}
		if conds.TestPlanID != nil {
			if _, err := uuid.Parse(conds.GetTestPlanID().GetValue()); err != nil {
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
