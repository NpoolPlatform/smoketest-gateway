package plantestcase

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *string
	TestPlanID  *string
	TestCaseID  *string
	TestUserID  *string
	Input       *string
	Output      *string
	Result      *pb.TestCaseResult
	Description *string
	Index       *uint32
	RunDuration *uint32
	Offset      int32
	Limit       int32
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
		if id == nil {
			return nil
		}
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

func WithTestUserID(userID, appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			return nil
		}
		if userID != nil && appID == nil {
			return fmt.Errorf("app id is empty")
		}
		if _, err := uuid.Parse(*userID); err != nil {
			return err
		}
		if _, err := uuid.Parse(*appID); err != nil {
			return err
		}
		// TODO: Query User By AppID & UserID
		h.TestUserID = userID
		return nil
	}
}

func WithInput(input *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if input == nil {
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*input), &r); err != nil {
			return err
		}
		h.Input = input
		return nil
	}
}

func WithOutput(output *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if output == nil {
			return nil
		}
		h.Output = output
		return nil
	}
}

func WithResult(result *pb.TestCaseResult) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			return nil
		}
		switch *result {
		case pb.TestCaseResult_Passed:
		case pb.TestCaseResult_Failed:
		case pb.TestCaseResult_Skipped:
		default:
			return fmt.Errorf("invalid result")
		}

		h.Result = result
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
