package biz

import (
	"context"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
)

type DemoRepo interface {
	Create(ctx context.Context, um *micro.UserContextMeta, request *pb.CreateRequest) error
	FindList(ctx context.Context, um *micro.UserContextMeta, request *pb.FindListRequest) *pb.List
	FindById(ctx context.Context, id string) (*pb.Demo, error)
	Update(ctx context.Context, um *micro.UserContextMeta, request *pb.UpdateRequest) error
	DeleteById(ctx context.Context, id string) error
}

type DemoUseCase struct {
	repo DemoRepo
}

func NewDemoUseCase(repo DemoRepo) *DemoUseCase {
	return &DemoUseCase{repo: repo}
}

func (uc *DemoUseCase) Create(ctx context.Context, um *micro.UserContextMeta, request *pb.CreateRequest) error {
	return uc.repo.Create(ctx, um, request)
}

func (uc *DemoUseCase) FindList(ctx context.Context, um *micro.UserContextMeta, request *pb.FindListRequest) *pb.List {
	return uc.repo.FindList(ctx, um, request)
}

func (uc *DemoUseCase) FindById(ctx context.Context, id string) (*pb.Demo, error) {
	return uc.repo.FindById(ctx, id)
}

func (uc *DemoUseCase) Update(ctx context.Context, um *micro.UserContextMeta, request *pb.UpdateRequest) error {
	return uc.repo.Update(ctx, um, request)
}

func (uc *DemoUseCase) DeleteById(ctx context.Context, id string) error {
	return uc.repo.DeleteById(ctx, id)
}
