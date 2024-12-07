package service

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
	"time"
)

func SaveDTO(src *pb.SaveRequest) *biz.Demo {
	var dst biz.Demo
	dst.Title = src.Title
	dst.Description = src.Description
	dst.Content = src.Content
	dst.Status = src.Status
	dst.Sort = src.Sort
	return &dst
}

func GetDTO(src *biz.Demo) *pb.Demo {
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
