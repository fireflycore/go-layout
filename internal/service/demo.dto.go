package service

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
)

func CreateDTO(src *pb.CreateRequest) *biz.Demo {
	var dst biz.Demo
	dst.Title = src.Title
	dst.Description = src.Description
	dst.Content = src.Content
	dst.Status = src.Status
	dst.Sort = src.Sort
	return &dst
}
