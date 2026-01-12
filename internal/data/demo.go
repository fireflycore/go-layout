package data

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/data/entity"

	micro "github.com/fireflycore/go-micro/rpc"
	"github.com/fireflycore/gormx/scope"
)

type demoRepo struct {
	data *Data
}

func NewDemoRepo(data *Data) repo.DemoRepo {
	return &demoRepo{
		data: data,
	}
}

func (uc *demoRepo) CreateDemo(ctx context.Context, row *entity.Demo) error {
	return uc.data.db.WithContext(ctx).Create(row).Error
}

func (uc *demoRepo) GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) *pb.DemoList {
	var raw pb.DemoList

	sql := uc.data.db.WithContext(ctx).Model(&entity.Demo{})
	sql.Where("user_id = ?", um.UserId)
	sql.Where("app_id = ?", um.AppId)
	if len(request.SearchKey) != 0 {
		sk := "%" + request.SearchKey + "%"
		sql.Where("name LIKE ?", sk)
	}
	sql.Count(&raw.Total)
	sql.Scopes(scope.WithPagination(request.Page, request.PageSize))
	sql.Find(&raw.List)

	return &raw
}

func (uc *demoRepo) GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error) {
	var row pb.Demo
	if res := uc.data.db.WithContext(ctx).Model(&entity.Demo{}).Where("id = ?", id).Find(&row); res.Error != nil {
		return nil, res.Error
	}
	return &row, nil
}

func (uc *demoRepo) UpdateDemo(ctx context.Context, id string, row map[string]interface{}) error {
	return uc.data.db.WithContext(ctx).Model(&entity.Demo{}).Where("id = ?", id).Updates(&row).Error
}

func (uc *demoRepo) DeleteDemo(ctx context.Context, id string) error {
	uc.data.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Demo{})
	return nil
}

func (uc *demoRepo) GetDemoCount(ctx context.Context, status *uint32) int64 {
	var count int64

	sql := uc.data.db.WithContext(ctx).Model(&entity.Demo{})
	if status != nil {
		sql.Where("status = ?", *status)
	}
	sql.Count(&count)

	return count
}
