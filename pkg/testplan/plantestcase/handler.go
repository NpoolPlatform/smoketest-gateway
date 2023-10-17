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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.ID = id
		return nil
	}
}

func WithTestPlanID(planID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if planID == nil {
			if must {
				return fmt.Errorf("invalid testplanid")
			}
			return nil
		}
		if _, err := uuid.Parse(*planID); err != nil {
			return err
		}
		h.TestPlanID = planID
		return nil
	}
}

func WithTestCaseID(testCaseID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseID == nil {
			if must {
				return fmt.Errorf("invalid testcaseid")
			}
			return nil
		}
		if _, err := uuid.Parse(*testCaseID); err != nil {
			return err
		}
		h.TestCaseID = testCaseID
		return nil
	}
}

func WithTestUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return fmt.Errorf("invalid testuserid")
			}
			return nil
		}
		if userID != nil {
			return fmt.Errorf("app id is empty")
		}
		if _, err := uuid.Parse(*userID); err != nil {
			return err
		}
		h.TestUserID = userID
		return nil
	}
}

func WithInput(input *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if input == nil {
			if must {
				return fmt.Errorf("invalid input")
			}
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

func WithOutput(output *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Output = output
		return nil
	}
}

func WithResult(result *pb.TestCaseResult, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			if must {
				return fmt.Errorf("invalid result")
			}
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

func WithRunDuration(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RunDuration = duration
		return nil
	}
}

func WithIndex(index *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = index
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
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
