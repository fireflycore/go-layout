package data

import (
	"context"
	"github.com/lhdhtrc/func-go/object"
	"github.com/lhdhtrc/gorm/pkg"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
)

type demoRepo struct {
	data *Data
}

func NewDemoRepo(data *Data) biz.DemoRepo {
	return &demoRepo{
		data: data,
	}
}

func (uc *demoRepo) Save(ctx context.Context, row *biz.Demo) error {
	uc.data.Mysql.WithContext(ctx).Create(&row)
	return nil
}

func (uc *demoRepo) Update(ctx context.Context, request *pb.UpdateRequest) error {
	var row biz.Demo
	if res := uc.data.Mysql.WithContext(ctx).Find(&row); res.Error != nil {
		return res.Error
	}

	update := object.FilterChangeValue(&row, request, []string{"Id", "Type"})
	if len(update) != 0 {
		uc.data.Mysql.WithContext(ctx).Updates(&row)
	}

	return nil
}

func (uc *demoRepo) FindById(ctx context.Context, id string) (*biz.Demo, error) {
	var row biz.Demo
	uc.data.Mysql.WithContext(ctx).Where("id = ?", id).Find(&row)
	return &row, nil
}

func (uc *demoRepo) FindList(ctx context.Context, query *pb.GetListRequest) (int64, []*pb.Demo) {
	sql := uc.data.Mysql.WithContext(ctx).Model(&biz.Demo{})

	var total int64
	var list []*pb.Demo
	sql.Count(&total)

	gorm.WithPagingFilter(sql, query.Page, query.PageSize)
	sql.Find(&list)

	return total, list
}

func (uc *demoRepo) DeleteById(ctx context.Context, id string) error {
	uc.data.Mysql.WithContext(ctx).Where("id = ?", id).Delete(&biz.Demo{})
	return nil
}
