package biz

import (
	"context"
	gorm "github.com/lhdhtrc/gorm/pkg"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
)

type Demo struct {
	gorm.TableUUID

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`

	AppId     string `json:"app_id" gorm:"size:36;index;"`
	AccountId string `json:"account_id" gorm:"size:36;index;"`
}

func (Demo) Table() string {
	return "demo"
}

type DemoRepo interface {
	Create(ctx context.Context, row *Demo) error
	Update(ctx context.Context, row *pb.UpdateRequest) error
	FindById(ctx context.Context, id string) (*pb.Demo, error)
	FindList(ctx context.Context, request *pb.FindListRequest) *pb.List

	DeleteById(ctx context.Context, id string)
	DeleteByIds(ctx context.Context, ids []string)
	DeleteByAppId(ctx context.Context, appId string)
	DeleteByAccountId(ctx context.Context, accountId string)
}

type DemoUseCase struct {
	repo DemoRepo
}

func NewDemoUseCase(repo DemoRepo) *DemoUseCase {
	return &DemoUseCase{repo: repo}
}

func (uc *DemoUseCase) Create(ctx context.Context, row *Demo) error {
	return uc.repo.Create(ctx, row)
}

func (uc *DemoUseCase) FindById(ctx context.Context, id string) (*pb.Demo, error) {
	return uc.repo.FindById(ctx, id)
}

func (uc *DemoUseCase) FindList(ctx context.Context, request *pb.FindListRequest) *pb.List {
	return uc.repo.FindList(ctx, request)
}

func (uc *DemoUseCase) Update(ctx context.Context, row *pb.UpdateRequest) error {
	return uc.repo.Update(ctx, row)
}

func (uc *DemoUseCase) DeleteById(ctx context.Context, id string) {
	uc.repo.DeleteById(ctx, id)
}
