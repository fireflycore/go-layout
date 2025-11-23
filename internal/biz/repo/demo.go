package repo

import (
	"context"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"
)

type DemoRepo interface {
	CreateDemo(ctx context.Context, row *entity.Demo) error
	GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) ([]*entity.Demo, int64)
	GetDemoInfo(ctx context.Context, id string) (*entity.Demo, error)
	UpdateDemo(ctx context.Context, id string, row map[string]interface{}) error
	DeleteDemo(ctx context.Context, id string) error
}
