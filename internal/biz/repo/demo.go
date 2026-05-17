package repo

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"

	"github.com/fireflycore/go-micro/service"
)

// DemoRepo 定义 Demo 仓储接口。
type DemoRepo interface {
	// CreateDemo 创建 Demo 记录。
	CreateDemo(ctx context.Context, row *entity.Demo) error
	// GetDemoList 按服务上下文边界分页查询 Demo 列表。
	GetDemoList(ctx context.Context, um *service.Context, request *pb.GetDemoListRequest) *pb.GetDemoListResponse
	// GetDemoInfo 按主键查询单个 Demo。
	GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error)
	// UpdateDemo 按主键更新 Demo 的指定字段。
	UpdateDemo(ctx context.Context, id string, row map[string]any) error
	// DeleteDemo 按主键删除 Demo。
	DeleteDemo(ctx context.Context, id string) error
	// GetDemoCount 按条件统计 Demo 数量。
	GetDemoCount(ctx context.Context, status *uint32) int64
}
