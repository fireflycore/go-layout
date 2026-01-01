package repo

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"

	micro "github.com/lhdhtrc/micro-go/pkg/core"
)

type DemoRepo interface {
	CreateDemo(ctx context.Context, row *entity.Demo) error
	GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) *pb.DemoList
	GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error)
	UpdateDemo(ctx context.Context, id string, row map[string]interface{}) error
	DeleteDemo(ctx context.Context, id string) error
	GetCount(ctx context.Context, status *uint32) (int64, error)
}
