package biz

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/model"
)

type DemoRepo interface {
	Save(ctx context.Context, row *model.DemoEntity) error
	Update(ctx context.Context, row *pb.UpdateRequest) error
	FindById(ctx context.Context, id string) (*model.DemoEntity, error)
	FindList(ctx context.Context, query *pb.GetListRequest) (int64, []*pb.Demo)
	DeleteById(ctx context.Context, id string) error
}

type DemoUseCase struct {
	repo DemoRepo
}

func NewDemoUseCase(repo DemoRepo) *DemoUseCase {
	return &DemoUseCase{repo: repo}
}

func (uc *DemoUseCase) Save(ctx context.Context, row *model.DemoEntity) error {
	return uc.repo.Save(ctx, row)
}

func (uc *DemoUseCase) FindById(ctx context.Context, id string) (*model.DemoEntity, error) {
	return uc.repo.FindById(ctx, id)
}

func (uc *DemoUseCase) FindList(ctx context.Context, query *pb.GetListRequest) (int64, []*pb.Demo) {
	return uc.repo.FindList(ctx, query)
}

func (uc *DemoUseCase) Update(ctx context.Context, row *pb.UpdateRequest) error {
	return uc.repo.Update(ctx, row)
}

func (uc *DemoUseCase) DeleteById(ctx context.Context, id string) error {
	return uc.repo.DeleteById(ctx, id)
}
