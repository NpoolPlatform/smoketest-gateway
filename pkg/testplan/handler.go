package testplan

import (
	"context"
	"fmt"
	"time"

	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *string
	Name        *string
	State       *pb.TestPlanState
	CreatedBy   *string
	Executor    *string
	Deadline    *uint32
	Fails       *uint32
	Skips       *uint32
	Passes      *uint32
	Result      *pb.TestResultState
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
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

func WithCreatedBy(createdBy *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if createdBy == nil {
			if must {
				return fmt.Errorf("ivnalid createdby")
			}
			return nil
		}
		if _, err := uuid.Parse(*createdBy); err != nil {
			return err
		}
		h.CreatedBy = createdBy
		return nil
	}
}

func WithDeadline(deadline *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deadline == nil {
			if must {
				return fmt.Errorf("invalid deadline")
			}
			return nil
		}
		if *deadline <= uint32(time.Now().Unix()) {
			return fmt.Errorf("deadline less than current time")
		}

		h.Deadline = deadline
		return nil
	}
}

func WithExecutor(executor *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if executor == nil {
			if must {
				return fmt.Errorf("invalid executor")
			}
			return nil
		}
		if _, err := uuid.Parse(*executor); err != nil {
			return err
		}
		h.Executor = executor
		return nil
	}
}

func WithState(state *pb.TestPlanState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		switch *state {
		case pb.TestPlanState_WaitStart:
		case pb.TestPlanState_InProgress:
		case pb.TestPlanState_Finished:
		case pb.TestPlanState_Overdue:
		default:
			return fmt.Errorf("plan state %v invalid", *state)
		}
		h.State = state
		return nil
	}
}

func WithResult(result *pb.TestResultState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			if must {
				return fmt.Errorf("invalid result")
			}
			return nil
		}
		switch *result {
		case pb.TestResultState_Failed:
		case pb.TestResultState_Passed:
		default:
			return fmt.Errorf("plan result %v invalid", *result)
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

func WithFails(fails *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Fails = fails
		return nil
	}
}

func WithPasses(passes *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Passes = passes
		return nil
	}
}

func WithSkips(skips *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Skips = skips
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
