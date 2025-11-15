package service

import (
	"buf.build/go/protovalidate"
	"context"
	"errors"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
)

type DemoService struct {
	pb.UnimplementedDemoServiceServer

	uc *biz.DemoUseCase
}

func NewDemoService(uc *biz.DemoUseCase) *DemoService {
	return &DemoService{uc: uc}
}

func (srv *DemoService) CreateDemo(ctx context.Context, request *pb.CreateDemoRequest) (*pb.CreateDemoResponse, error) {
	result := &pb.CreateDemoResponse{
		Code:    200,
		Message: "success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	md, _ := metadata.FromIncomingContext(ctx)
	um, err := micro.ParseUserContextMeta(md)
	if err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, err
	}

	if err = srv.uc.CreateDemo(ctx, um, request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	return result, nil
}

func (srv *DemoService) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	result := &pb.UpdateResponse{
		Code:    200,
		Message: "update demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	md, _ := metadata.FromIncomingContext(ctx)
	um, err := micro.ParseUserContextMeta(md)
	if err != nil {
		result.Code = 400
		result.Message = "查询失败"
		return result, err
	}

	if err = srv.uc.Update(ctx, um, request); err != nil {
		result.Code = 400
		result.Message = "更新失败"
		return result, err
	}

	return result, nil
}

func (srv *DemoService) FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdResponse, error) {
	result := &pb.FindByIdResponse{
		Code:    200,
		Message: "get demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	row, err := srv.uc.FindById(ctx, request.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result.Code = 404
		result.Message = "resource not found"
		return result, err
	}
	result.Data = row

	return result, nil
}

func (srv *DemoService) FindList(ctx context.Context, request *pb.FindListRequest) (*pb.FindListResponse, error) {
	result := &pb.FindListResponse{
		Code:    200,
		Message: "get demo list success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	md, _ := metadata.FromIncomingContext(ctx)
	um, err := micro.ParseUserContextMeta(md)
	if err != nil {
		result.Code = 400
		result.Message = "查询失败"
		return result, err
	}

	result.Data = srv.uc.FindList(ctx, um, request)

	return result, nil
}

func (srv *DemoService) DeleteById(ctx context.Context, request *pb.DeleteByIdRequest) (*pb.DeleteByIdResponse, error) {
	result := &pb.DeleteByIdResponse{
		Code:    200,
		Message: "delete demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	if err := srv.uc.DeleteById(ctx, request.Id); err != nil {
		result.Code = 400
		result.Message = "删除失败"
		return result, err
	}

	return result, nil
}
