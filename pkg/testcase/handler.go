package testcase

import (
	"context"
	"encoding/json"
	"fmt"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *string
	Name          *string
	Description   *string
	ModuleID      *string
	ApiID         *string //nolint
	Input         *string
	InputDesc     *string
	Expectation   *string
	OutputDesc    *string
	TestCaseType  *pb.TestCaseType
	TestCaseClass *pb.TestCaseClass
	Deprecated    *bool
	Offset        int32
	Limit         int32
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

func WithModuleID(moduleID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if moduleID == nil {
			if must {
				return fmt.Errorf("invalid moduleid")
			}
			return nil
		}
		if _, err := uuid.Parse(*moduleID); err != nil {
			return err
		}
		h.ModuleID = moduleID
		return nil
	}
}

//nolint
func WithApiID(apiID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if apiID == nil {
			if must {
				return fmt.Errorf("invalid apiid")
			}
			return nil
		}
		exist, err := apicli.ExistAPI(ctx, *apiID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid apiid")
		}

		h.ApiID = apiID
		return nil
	}
}

func WithExpectation(expectation *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if expectation == nil {
			if must {
				return fmt.Errorf("invalid expectation")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*expectation), &r); err != nil {
			return err
		}
		h.Expectation = expectation
		return nil
	}
}

func WithOutputDesc(outputDesc *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if outputDesc == nil {
			if must {
				return fmt.Errorf("invalid outputdesc")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*outputDesc), &r); err != nil {
			return err
		}
		h.OutputDesc = outputDesc
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

func WithInputDesc(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			if must {
				return fmt.Errorf("invalid inputdesc")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*description), &r); err != nil {
			return err
		}
		h.InputDesc = description
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Description = description
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
		const leastNameLen = 5
		if len(*name) < leastNameLen {
			return fmt.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}

func WithDeprecated(deprecated *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Deprecated = deprecated
		return nil
	}
}

func WithTestCaseType(testCaseType *pb.TestCaseType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseType == nil {
			if must {
				return fmt.Errorf("invalid testcasetype")
			}
			return nil
		}
		switch *testCaseType {
		case pb.TestCaseType_Manual:
		case pb.TestCaseType_Automatic:
		default:
			return fmt.Errorf("invalid testcase type")
		}
		h.TestCaseType = testCaseType
		return nil
	}
}

func WithTestCaseClass(testCaseClass *pb.TestCaseClass, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseClass == nil {
			if must {
				return fmt.Errorf("invalid testcaseclass")
			}
			return nil
		}
		switch *testCaseClass {
		case pb.TestCaseClass_Functionality:
		case pb.TestCaseClass_Interface:
		default:
			return fmt.Errorf("invalid testcaseclass")
		}
		h.TestCaseClass = testCaseClass
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
