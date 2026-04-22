package repo

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"

	"github.com/fireflycore/go-micro/invocation"
)

type DemoRepo interface {
	CreateDemo(ctx context.Context, row *entity.Demo) error
	GetDemoList(ctx context.Context, um *invocation.UserContextMeta, request *pb.GetDemoListRequest) *pb.DemoList
	GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error)
	UpdateDemo(ctx context.Context, id string, row map[string]any) error
	DeleteDemo(ctx context.Context, id string) error
	GetDemoCount(ctx context.Context, status *uint32) int64
}
