package biz

import (
	"context"
	"github.com/lhdhtrc/func-go/object"
	gorme "github.com/lhdhtrc/gorm/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/depend/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/dto/convert"
)

type DemoUseCase struct {
	repo repo.DemoRepo

	dto convert.DemoConverter
}

func NewDemoUseCase(repo repo.DemoRepo, dto convert.DemoConverter) *DemoUseCase {
	return &DemoUseCase{
		repo: repo,

		dto: dto,
	}
}

func (uc *DemoUseCase) CreateDemo(ctx context.Context, um *micro.UserContextMeta, request *pb.CreateDemoRequest) error {
	row := uc.dto.ToCreate(request)
	row.AppId = gorme.ParseUUID(um.AppId)
	row.UserId = gorme.ParseUUID(um.UserId)

	return uc.repo.CreateDemo(ctx, row)
}

func (uc *DemoUseCase) GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) *pb.DemoList {
	list, total := uc.repo.GetDemoList(ctx, um, request)

	return &pb.DemoList{
		Total: total,
		List:  uc.dto.ToRaw(list),
	}
}

func (uc *DemoUseCase) GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error) {
	row, err := uc.repo.GetDemoInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return uc.dto.ToRow(row), nil
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
