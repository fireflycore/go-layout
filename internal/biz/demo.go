package biz

import (
	"context"
	"github.com/lhdhtrc/func-go/object"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/convert"
	"go-layout/internal/biz/repo"
)

type DemoUseCase struct {
	repo repo.DemoRepo

	dto convert.DemoConvert
}

func NewDemoUseCase(repo repo.DemoRepo, dto convert.DemoConvert) *DemoUseCase {
	return &DemoUseCase{
		repo: repo,

		dto: dto,
	}
}

func (uc *DemoUseCase) CreateDemo(ctx context.Context, um *micro.UserContextMeta, request *pb.CreateDemoRequest) error {
	row := uc.dto.ToCreate(request)
	row.AppId = um.AppId
	row.UserId = um.UserId
	row.TenantId = um.TenantId

	return uc.repo.CreateDemo(ctx, row)
}

func (uc *DemoUseCase) GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) *pb.DemoList {
	return uc.repo.GetDemoList(ctx, um, request)
}

func (uc *DemoUseCase) GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error) {
	return uc.repo.GetDemoInfo(ctx, id)
}

func (uc *DemoUseCase) UpdateDemo(ctx context.Context, _ *micro.UserContextMeta, request *pb.UpdateDemoRequest) error {
	row, err := uc.repo.GetDemoInfo(ctx, request.Id)
	if err != nil {
		return err
	}

	updates := object.FilterChangeValue(row, request, []string{})
	if len(updates) == 0 {
		return nil
	}

	return uc.repo.UpdateDemo(ctx, request.Id, updates)
}

func (uc *DemoUseCase) DeleteDemo(ctx context.Context, id string) error {
	return uc.repo.DeleteDemo(ctx, id)
}
