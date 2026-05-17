package service

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"

	"buf.build/go/protovalidate"
)

// DemoService 实现 Demo 相关的 gRPC Service。
type DemoService struct {
	pb.UnimplementedDemoServiceServer

	uc *biz.DemoUseCase
}

// NewDemoService 创建 Demo gRPC Service 入口。
func NewDemoService(uc *biz.DemoUseCase) *DemoService {
	return &DemoService{uc: uc}
}

// CreateDemo 处理 Demo 创建请求。
func (srv *DemoService) CreateDemo(ctx context.Context, request *pb.CreateDemoRequest) (*pb.CreateDemoResponse, error) {
	// 入口统一先做 protovalidate 校验，避免 Service 层重复实现字段规则。
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	// Service 只做入口编排，真正的创建逻辑由 UseCase 承担。
	if err := srv.uc.CreateDemo(ctx, request); err != nil {
		return nil, err
	}

	return &pb.CreateDemoResponse{}, nil
}

// GetDemoList 处理 Demo 列表查询请求。
func (srv *DemoService) GetDemoList(ctx context.Context, request *pb.GetDemoListRequest) (*pb.GetDemoListResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	return srv.uc.GetDemoList(ctx, request), nil
}

// GetDemoInfo 处理单个 Demo 详情查询请求。
func (srv *DemoService) GetDemoInfo(ctx context.Context, request *pb.GetDemoInfoRequest) (*pb.GetDemoInfoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	// 详情读取完成后由 Service 组装标准响应壳。
	row, err := srv.uc.GetDemoInfo(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDemoInfoResponse{Data: row}, nil
}

// UpdateDemo 处理 Demo 更新请求。
func (srv *DemoService) UpdateDemo(ctx context.Context, request *pb.UpdateDemoRequest) (*pb.UpdateDemoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	// 字段差异计算和最终更新判断都放在 UseCase 中完成。
	if err := srv.uc.UpdateDemo(ctx, request); err != nil {
		return nil, err
	}

	return &pb.UpdateDemoResponse{}, nil
}

// DeleteDemo 处理 Demo 删除请求。
func (srv *DemoService) DeleteDemo(ctx context.Context, request *pb.DeleteDemoRequest) (*pb.DeleteDemoResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	// 删除动作只透传上下文与主键，不在 Service 层叠加额外业务逻辑。
	if err := srv.uc.DeleteDemo(ctx, request.Id); err != nil {
		return nil, err
	}

	return &pb.DeleteDemoResponse{}, nil
}

// GetDemoCount 处理 Demo 数量统计请求。
func (srv *DemoService) GetDemoCount(ctx context.Context, request *pb.GetDemoCountRequest) (*pb.GetDemoCountResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, err
	}

	return &pb.GetDemoCountResponse{Data: srv.uc.GetDemoCount(ctx, request.Status)}, nil
}
