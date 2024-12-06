package service

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
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

func (s *DemoService) Save(ctx context.Context, request *pb.SaveRequest) (*pb.SaveResponse, error) {
	result := &pb.SaveResponse{
		Code:    200,
		Message: "save demo success",
	}

	md, _ := metadata.FromIncomingContext(ctx)

	appId := md.Get("app-id")
	accountId := md.Get("account-id")

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	row := SaveDTO(request)
	row.AppId = appId[0]
	row.AccountId = accountId[0]

	err := s.uc.Save(ctx, row)

	return result, err
}

func (s *DemoService) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	result := &pb.UpdateResponse{
		Code:    200,
		Message: "update demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	if err := s.uc.Update(ctx, request); errors.Is(err, gorm.ErrRecordNotFound) {
		result.Code = 404
		result.Message = "resource not found"
		return nil, err
	}

	return result, nil
}

func (s *DemoService) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	result := &pb.GetResponse{
		Code:    200,
		Message: "get demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	row, err := s.uc.FindById(ctx, request.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result.Code = 404
		result.Message = "resource not found"
		return result, err
	}
	result.Data = GetDTO(row)

	return result, nil
}

func (s *DemoService) GetList(ctx context.Context, request *pb.GetListRequest) (*pb.GetListResponse, error) {
	result := &pb.GetListResponse{
		Code:    200,
		Message: "get demo list success",
		Data:    &pb.GetList{},
	}

	total, list := s.uc.FindList(ctx, request)
	result.Data = &pb.GetList{
		Total: total,
		List:  list,
	}

	return result, nil
}

func (s *DemoService) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result := &pb.DeleteResponse{
		Code:    200,
		Message: "delete demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	err := s.uc.DeleteById(ctx, request.Id)

	return result, err
}
