package data

import (
	"context"
	"github.com/lhdhtrc/func-go/object"
	gorms "github.com/lhdhtrc/gorm/pkg"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
	"go-layout/model"
)

type demoRepo struct {
	data *Data
}

func NewDemoRepo(data *Data) biz.DemoRepo {
	return &demoRepo{
		data: data,
	}
}

func (uc *demoRepo) Save(ctx context.Context, row *model.DemoEntity) error {
	uc.data.Mysql.WithContext(ctx).Create(&row)
	return nil
}

func (uc *demoRepo) Update(ctx context.Context, request *pb.UpdateRequest) error {
	var row model.DemoEntity
	if res := uc.data.Mysql.WithContext(ctx).Find(&row); res.Error != nil {
		return res.Error
	}

	update := object.FilterChangeValue(&row, request, []string{"Id", "Type"})
	if len(update) != 0 {
		uc.data.Mysql.WithContext(ctx).Updates(&row)
	}

	return nil
}

func (uc *demoRepo) FindById(ctx context.Context, id string) (*model.DemoEntity, error) {
	var row model.DemoEntity
	uc.data.Mysql.WithContext(ctx).Where("id = ?", id).Find(&row)
	return &row, nil
}

func (uc *demoRepo) FindList(ctx context.Context, query *pb.GetListRequest) (int64, []*pb.Demo) {
	sql := uc.data.Mysql.WithContext(ctx).Model(&model.DemoEntity{})

	var total int64
	var list []*pb.Demo
	sql.Count(&total)

	gorms.WithPagingFilter(sql, query.Page, query.PageSize)
	sql.Find(&list)

	return total, list
}

func (uc *demoRepo) DeleteById(ctx context.Context, id string) error {
	uc.data.Mysql.WithContext(ctx).Where("id = ?", id).Delete(&model.DemoEntity{})
	return nil
}
