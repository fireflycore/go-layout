package service

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"

	"buf.build/go/protovalidate"
	micro "github.com/fireflycore/go-micro/rpc"
	"google.golang.org/grpc/metadata"
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
		return result, nil
	}

	if err = srv.uc.CreateDemo(ctx, um, request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	return result, nil
}

func (srv *DemoService) GetDemoList(ctx context.Context, request *pb.GetDemoListRequest) (*pb.GetDemoListResponse, error) {
	result := &pb.GetDemoListResponse{
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
		return result, nil
	}

	result.Data = srv.uc.GetDemoList(ctx, um, request)

	return result, nil
}

func (srv *DemoService) GetDemoInfo(ctx context.Context, request *pb.GetDemoInfoRequest) (*pb.GetDemoInfoResponse, error) {
	result := &pb.GetDemoInfoResponse{
		Code:    200,
		Message: "success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	row, err := srv.uc.GetDemoInfo(ctx, request.Id)
	if err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}
	result.Data = row

	return result, nil
}

func (srv *DemoService) UpdateDemo(ctx context.Context, request *pb.UpdateDemoRequest) (*pb.UpdateDemoResponse, error) {
	result := &pb.UpdateDemoResponse{
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
		return result, nil
	}

	if err = srv.uc.UpdateDemo(ctx, um, request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	return result, nil
}

func (srv *DemoService) DeleteDemo(ctx context.Context, request *pb.DeleteDemoRequest) (*pb.DeleteDemoResponse, error) {
	result := &pb.DeleteDemoResponse{
		Code:    200,
		Message: "success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	if err := srv.uc.DeleteDemo(ctx, request.Id); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	return result, nil
}

func (srv *DemoService) GetDemoCount(ctx context.Context, request *pb.GetDemoCountRequest) (*pb.GetDemoCountResponse, error) {
	result := &pb.GetDemoCountResponse{
		Code:    200,
		Message: "success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = err.Error()
		return result, nil
	}

	result.Data = srv.uc.GetDemoCount(ctx, request.Status)

	return result, nil
}
