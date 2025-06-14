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

func (uc *demoRepo) Create(ctx context.Context, row *biz.Demo) error {
	uc.data.db.WithContext(ctx).Create(&row)
	return nil
}

func (uc *demoRepo) Update(ctx context.Context, request *pb.UpdateRequest) error {
	var row biz.Demo
	if res := uc.data.db.WithContext(ctx).Find(&row); res.Error != nil {
		return res.Error
	}

	update := object.FilterChangeValue(&row, request, []string{"Id", "Type"})
	if len(update) != 0 {
		uc.data.db.WithContext(ctx).Updates(&row)
	}

	return nil
}

func (uc *demoRepo) FindById(ctx context.Context, id string) (*pb.Demo, error) {
	var row *pb.Demo
	if res := uc.data.db.WithContext(ctx).Model(&biz.Demo{}).Where("id = ?", id).Find(&row); res.Error != nil {
		return nil, res.Error
	}
	return row, nil
}

func (uc *demoRepo) FindList(ctx context.Context, request *pb.FindListRequest) *pb.List {
	var raw *pb.List

	sql := uc.data.db.WithContext(ctx).Model(&biz.Demo{})
	sql.Count(&raw.Total)

	gorm.UsePaging(sql, request.Page, request.PageSize)
	sql.Find(&raw.List)

	return raw
}

func (uc *demoRepo) DeleteById(ctx context.Context, id string) {
	uc.data.db.WithContext(ctx).Where("id = ?", id).Delete(&biz.Demo{})
}

func (uc *demoRepo) DeleteByIds(ctx context.Context, ids []string) {
	uc.data.db.WithContext(ctx).Where("id IN ?", ids).Delete(&biz.Demo{})
}

func (uc *demoRepo) DeleteByAppId(ctx context.Context, appId string) {
	uc.data.db.WithContext(ctx).Where("app_id = ?", appId).Delete(&biz.Demo{})
}

func (uc *demoRepo) DeleteByAccountId(ctx context.Context, accountId string) {
	uc.data.db.WithContext(ctx).Where("account_id = ?", accountId).Delete(&biz.Demo{})
}
