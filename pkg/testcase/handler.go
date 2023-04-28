package testcase

import (
	"context"
	"encoding/json"
	"fmt"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID           *string
	Name         *string
	Description  *string
	ModuleName   *string
	ModuleID     *string
	ApiID        *string //nolint
	Input        *string
	InputDesc    *string
	Expectation  *string
	TestCaseType *pb.TestCaseType
	Deprecated   *bool
	Offset       int32
	Limit        int32
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

func WithModuleID(moduleID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if moduleID == nil {
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
func WithApiID(apiID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*apiID); err != nil {
			return err
		}

		exist, err := apimwcli.ExistAPI(ctx, *apiID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid api id")
		}

		h.ApiID = apiID
		return nil
	}
}

func WithModuleName(moduleName *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if moduleName == nil {
			return nil
		}
		const leastModuleNameLen = 2
		if len(*moduleName) < leastModuleNameLen {
			return fmt.Errorf("invalid module name")
		}
		h.ModuleName = moduleName
		return nil
	}
}

func WithExpectation(expectation *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if expectation == nil {
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

func WithInputDesc(description *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
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

func WithDescription(description *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			return nil
		}
		h.Description = description
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
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

func WithDeprecated(deprecated *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deprecated == nil {
			return nil
		}
		h.Deprecated = deprecated
		return nil
	}
}

func WithTestCaseType(testCaseType *pb.TestCaseType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseType == nil {
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
