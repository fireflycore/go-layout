package biz

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/convert"
	"go-layout/internal/biz/repo"

	"github.com/fireflycore/go-micro/service"
)

// DemoUseCase 承载 Demo 相关业务逻辑。
type DemoUseCase struct {
	dto  convert.DemoConvert
	repo repo.DemoRepo
}

// NewDemoUseCase 创建 Demo UseCase。
func NewDemoUseCase(dto convert.DemoConvert, repo repo.DemoRepo) *DemoUseCase {
	return &DemoUseCase{
		dto:  dto,
		repo: repo,
	}
}

// CreateDemo 创建 Demo 数据。
func (uc *DemoUseCase) CreateDemo(ctx context.Context, request *pb.CreateDemoRequest) error {
	um := service.MustFromContext(ctx)

	// 先把请求转换成实体，再由上下文补齐租户、应用和用户边界。
	row := uc.dto.ToCreate(request)
	row.AppId = um.AppId
	row.UserId = um.UserId
	row.TenantId = um.TenantId

	return uc.repo.CreateDemo(ctx, row)
}

// GetDemoList 查询 Demo 列表。
func (uc *DemoUseCase) GetDemoList(ctx context.Context, request *pb.GetDemoListRequest) *pb.GetDemoListResponse {
	um := service.MustFromContext(ctx)

	return uc.repo.GetDemoList(ctx, um, request)
}

// GetDemoInfo 查询单个 Demo 详情。
func (uc *DemoUseCase) GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error) {
	return uc.repo.GetDemoInfo(ctx, id)
}

// UpdateDemo 更新 Demo 数据。
func (uc *DemoUseCase) UpdateDemo(ctx context.Context, request *pb.UpdateDemoRequest) error {
	// 先读取当前快照，再按字段逐项比较，避免无变更时产生无意义写入。
	row, err := uc.repo.GetDemoInfo(ctx, request.Id)
	if err != nil {
		return err
	}

	updates := make(map[string]interface{})
	if request.Title != "" && request.Title != row.Title {
		updates["title"] = request.Title
	}
	if request.Description != "" && request.Description != row.Description {
		updates["description"] = request.Description
	}
	if request.Content != "" && request.Content != row.Content {
		updates["content"] = request.Content
	}
	if request.Status != 0 && request.Status != row.Status {
		updates["status"] = request.Status
	}
	if request.Sort != 0 && request.Sort != row.Sort {
		updates["sort"] = request.Sort
	}
	// 没有实际差异时直接返回，保持更新链路幂等。
	if len(updates) == 0 {
		return nil
	}

	return uc.repo.UpdateDemo(ctx, request.Id, updates)
}

// DeleteDemo 删除 Demo 数据。
func (uc *DemoUseCase) DeleteDemo(ctx context.Context, id string) error {
	return uc.repo.DeleteDemo(ctx, id)
}

// GetDemoCount 统计 Demo 数量。
func (uc *DemoUseCase) GetDemoCount(ctx context.Context, status *uint32) int64 {
	return uc.repo.GetDemoCount(ctx, status)
}
