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
	Save(ctx context.Context, row *Demo) error
	Update(ctx context.Context, row *pb.UpdateRequest) error
	FindById(ctx context.Context, id string) (*Demo, error)
	FindList(ctx context.Context, query *pb.GetListRequest) (int64, []*pb.Demo)
	DeleteById(ctx context.Context, id string) error
}

type DemoUseCase struct {
	repo DemoRepo
}

func NewDemoUseCase(repo DemoRepo) *DemoUseCase {
	return &DemoUseCase{repo: repo}
}

func (uc *DemoUseCase) Save(ctx context.Context, row *Demo) error {
	return uc.repo.Save(ctx, row)
}

func (uc *DemoUseCase) FindById(ctx context.Context, id string) (*Demo, error) {
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
