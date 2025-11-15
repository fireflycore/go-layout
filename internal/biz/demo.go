package biz

import (
	"context"
	gorme "github.com/lhdhtrc/gorm/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
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
