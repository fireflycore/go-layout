package service

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
	"go-layout/internal/service/dto"
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

func (s *DemoService) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	result := &pb.CreateResponse{
		Code:    200,
		Message: "create demo success",
	}

	md, _ := metadata.FromIncomingContext(ctx)

	appId := md.Get("app-id")
	accountId := md.Get("account-id")

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	row := dto.CreateDTO(request)
	row.AppId = appId[0]
	row.AccountId = accountId[0]

	err := s.uc.Create(ctx, row)

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

func (s *DemoService) FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdResponse, error) {
	result := &pb.FindByIdResponse{
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
	result.Data = row

	return result, nil
}

func (s *DemoService) FindList(ctx context.Context, request *pb.FindListRequest) (*pb.FindListResponse, error) {
	result := &pb.FindListResponse{
		Code:    200,
		Message: "get demo list success",
	}

	result.Data = s.uc.FindList(ctx, request)

	return result, nil
}

func (s *DemoService) DeleteById(ctx context.Context, request *pb.DeleteByIdRequest) (*pb.DeleteByIdResponse, error) {
	result := &pb.DeleteByIdResponse{
		Code:    200,
		Message: "delete demo success",
	}

	if err := protovalidate.Validate(request); err != nil {
		result.Code = 400
		result.Message = "missing necessary params"
		return result, err
	}

	s.uc.DeleteById(ctx, request.Id)

	return result, nil
}
