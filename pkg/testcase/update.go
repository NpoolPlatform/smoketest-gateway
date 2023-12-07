package testcase

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/smoketest/gw/v1/testcase"
	pb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	cli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
)

func (h *Handler) UpdateTestCase(ctx context.Context) (*npool.TestCase, error) {
	info, err := cli.GetTestCaseOnly(ctx, &pb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	_, err = cli.UpdateTestCase(ctx, &pb.TestCaseReq{
		ID:            h.ID,
		Name:          h.Name,
		Description:   h.Description,
		Input:         h.Input,
		InputDesc:     h.InputDesc,
		Expectation:   h.Expectation,
		OutputDesc:    h.OutputDesc,
		TestCaseType:  h.TestCaseType,
		TestCaseClass: h.TestCaseClass,
		Deprecated:    h.Deprecated,
	},
	)
	if err != nil {
		return nil, err
	}
	return h.GetTestCase(ctx)
}
