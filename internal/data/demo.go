package data

import (
	"context"
	"errors"
	"github.com/lhdhtrc/func-go/object"
	"github.com/lhdhtrc/gorm/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
	"go-layout/internal/data/entity"
)

type demoRepo struct {
	data *Data
}

func NewDemoRepo(data *Data) biz.DemoRepo {
	return &demoRepo{
		data: data,
	}
}

func (uc *demoRepo) CreateDemo(ctx context.Context, row *entity.Demo) error {
	return uc.data.db.WithContext(ctx).Create(row).Error
}

func (uc *demoRepo) GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) ([]*entity.Demo, int64) {
	var (
		list  []*entity.Demo
		total int64
	)

	sql := uc.data.db.WithContext(ctx).Model(&entity.Demo{})
	sql.Where("user_id = ?", um.UserId)
	sql.Where("app_id = ?", um.AppId)
	if len(request.SearchKey) != 0 {
		sk := "%" + request.SearchKey + "%"
		sql.Where("name LIKE ?", sk)
	}
	sql.Count(&total)
	gorm.WithPaging(sql, request.Page, request.PageSize)
	sql.Find(&list)

	return list, total
}

func (uc *demoRepo) GetDemoInfo(ctx context.Context, id string) (*entity.Demo, error) {
	var row entity.Demo
	if res := uc.data.db.WithContext(ctx).Where("id = ?", id).Find(&row); res.Error != nil {
		return nil, res.Error
	}
	return &row, nil
}

func (uc *demoRepo) UpdateDemo(ctx context.Context, id string, row map[string]interface{}) error {
	return uc.data.db.WithContext(ctx).Model(&entity.Demo{}).Where("id = ?", id).Updates(&row).Error
}
