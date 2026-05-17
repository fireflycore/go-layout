package data

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/data/entity"

	"github.com/fireflycore/go-micro/service"
	"github.com/fireflycore/gormx/scope"
)

type demoRepo struct {
	data *Data
}

// NewDemoRepo 创建 Demo Repo 实现。
func NewDemoRepo(data *Data) repo.DemoRepo {
	return &demoRepo{
		data: data,
	}
}

// CreateDemo 创建 Demo 记录。
func (uc *demoRepo) CreateDemo(ctx context.Context, row *entity.Demo) error {
	return uc.data.db.WithContext(ctx).Create(row).Error
}

// GetDemoList 按服务上下文边界分页查询 Demo 列表。
func (uc *demoRepo) GetDemoList(ctx context.Context, sc *service.Context, request *pb.GetDemoListRequest) *pb.GetDemoListResponse {
	var raw pb.GetDemoListResponse

	sql := uc.data.db.WithContext(ctx).Model(&entity.Demo{})
	// Demo 列表按当前用户和应用过滤，避免跨应用或跨用户读取示例数据。
	sql.Where("user_id = ?", sc.UserId)
	sql.Where("app_id = ?", sc.AppId)
	// SearchKey 只做标题模糊检索，不参与身份判定。
	if len(request.SearchKey) != 0 {
		sk := "%" + request.SearchKey + "%"
		sql.Where("name LIKE ?", sk)
	}
	// 先 count 再分页，保证 total 反映完整结果集。
	sql.Count(&raw.Total)
	sql.Scopes(scope.WithPagination(request.Page, request.PageSize))
	sql.Find(&raw.List)

	return &raw
}

// GetDemoInfo 按主键查询单个 Demo。
func (uc *demoRepo) GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error) {
	var row pb.Demo
	// 详情查询显式按主键读取；错误语义直接向上返回。
	if res := uc.data.db.WithContext(ctx).Model(&entity.Demo{}).Where("id = ?", id).Find(&row); res.Error != nil {
		return nil, res.Error
	}
	return &row, nil
}

// UpdateDemo 按主键更新 Demo 的指定字段。
func (uc *demoRepo) UpdateDemo(ctx context.Context, id string, row map[string]interface{}) error {
	return uc.data.db.WithContext(ctx).Model(&entity.Demo{}).Where("id = ?", id).Updates(&row).Error
}

// DeleteDemo 按主键删除 Demo。
func (uc *demoRepo) DeleteDemo(ctx context.Context, id string) error {
	uc.data.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Demo{})
	return nil
}

// GetDemoCount 按条件统计 Demo 数量。
func (uc *demoRepo) GetDemoCount(ctx context.Context, status *uint32) int64 {
	var count int64

	sql := uc.data.db.WithContext(ctx).Model(&entity.Demo{})
	// status 为可选过滤条件，未传时返回全量计数。
	if status != nil {
		sql.Where("status = ?", *status)
	}
	sql.Count(&count)

	return count
}
