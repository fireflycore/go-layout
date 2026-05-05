package service

import (
	"context"
	"errors"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"

	"buf.build/go/protovalidate"
	servicectx "github.com/fireflycore/go-micro/service"
)

type DemoService struct {
	pb.UnimplementedDemoServiceServer

	uc *biz.DemoUseCase
}

func NewDemoService(uc *biz.DemoUseCase) *DemoService {
	return &DemoService{uc: uc}
}

func (srv *DemoService) CreateDemo(ctx context.Context, request *pb.CreateDemoRequest) (*pb.CreateDemoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	sc, err := requireServiceContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = srv.uc.CreateDemo(ctx, sc, request); err != nil {
		return nil, err
	}

	return &pb.CreateDemoResponse{}, nil
}

func (srv *DemoService) GetDemoList(ctx context.Context, request *pb.GetDemoListRequest) (*pb.GetDemoListResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	sc, err := requireServiceContext(ctx)
	if err != nil {
		return nil, err
	}

	return srv.uc.GetDemoList(ctx, sc, request), nil
}

func (srv *DemoService) GetDemoInfo(ctx context.Context, request *pb.GetDemoInfoRequest) (*pb.GetDemoInfoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	row, err := srv.uc.GetDemoInfo(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDemoInfoResponse{Data: row}, nil
}

func (srv *DemoService) UpdateDemo(ctx context.Context, request *pb.UpdateDemoRequest) (*pb.UpdateDemoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	sc, err := requireServiceContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = srv.uc.UpdateDemo(ctx, sc, request); err != nil {
		return nil, err
	}

	return &pb.UpdateDemoResponse{}, nil
}

func (srv *DemoService) DeleteDemo(ctx context.Context, request *pb.DeleteDemoRequest) (*pb.DeleteDemoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	if err := srv.uc.DeleteDemo(ctx, request.Id); err != nil {
		return nil, err
	}

	return &pb.DeleteDemoResponse{}, nil
}

func (srv *DemoService) GetDemoCount(ctx context.Context, request *pb.GetDemoCountRequest) (*pb.GetDemoCountResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	return &pb.GetDemoCountResponse{Data: srv.uc.GetDemoCount(ctx, request.Status)}, nil
}

func requireServiceContext(ctx context.Context) (*servicectx.Context, error) {
	sc, ok := servicectx.FromContext(ctx)
	if !ok {
		return nil, errors.New("service context not found")
	}
	return sc, nil
}
