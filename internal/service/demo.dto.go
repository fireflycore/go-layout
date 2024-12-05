package service

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/model"
	"time"
)

func SaveDTO(src *pb.SaveRequest) *model.DemoEntity {
	var dst model.DemoEntity
	dst.Title = src.Title
	dst.Description = src.Description
	dst.Content = src.Content
	dst.Status = src.Status
	dst.Sort = src.Sort
	return &dst
}

func GetDTO(src *model.DemoEntity) *pb.Demo {
	var dst pb.Demo
	dst.Id = src.ID
	dst.Title = src.Title
	dst.Description = src.Description
	dst.Content = src.Content
	dst.Status = src.Status
	dst.Sort = src.Sort
	dst.CreatedAt = src.CreatedAt.Format(time.DateTime)
	dst.UpdatedAt = src.UpdatedAt.Format(time.DateTime)
	dst.Author = &pb.Author{
		Id:   src.AccountId,
		Name: "",
	}
	return &dst
}
