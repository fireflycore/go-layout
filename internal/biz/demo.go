package biz

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/convert"
	"go-layout/internal/biz/repo"

	micro "github.com/fireflycore/go-micro/rpc"
)

type DemoUseCase struct {
	dto  convert.DemoConvert
	repo repo.DemoRepo
}

func NewDemoUseCase(dto convert.DemoConvert, repo repo.DemoRepo) *DemoUseCase {
	return &DemoUseCase{
		dto:  dto,
		repo: repo,
	}
}

func (uc *DemoUseCase) CreateDemo(ctx context.Context, um *micro.UserContextMeta, request *pb.CreateDemoRequest) error {
	row := uc.dto.ToCreate(request)
	row.AppId = um.AppId
	row.UserId = um.UserId
	row.TenantId = um.TenantId

	return uc.repo.CreateDemo(ctx, row)
}

func (uc *DemoUseCase) GetDemoList(ctx context.Context, um *micro.UserContextMeta, request *pb.GetDemoListRequest) *pb.DemoList {
	return uc.repo.GetDemoList(ctx, um, request)
}

func (uc *DemoUseCase) GetDemoInfo(ctx context.Context, id string) (*pb.Demo, error) {
	return uc.repo.GetDemoInfo(ctx, id)
}

func (uc *DemoUseCase) UpdateDemo(ctx context.Context, _ *micro.UserContextMeta, request *pb.UpdateDemoRequest) error {
	row, err := uc.repo.GetDemoInfo(ctx, request.Id)
	if err != nil {
		return err
	}

	updates := make(map[string]interface{})
	if request.Title != "" && request.Title != row.Title {
		updates["title"] = request.Title
	}
	if request.Description != "" && request.Description != row.Description {
		updates["description"] = request.Description
	}
	if request.Content != "" && request.Content != row.Content {
		updates["content"] = request.Content
	}
	if request.Status != 0 && request.Status != row.Status {
		updates["status"] = request.Status
	}
	if request.Sort != 0 && request.Sort != row.Sort {
		updates["sort"] = request.Sort
	}
	if len(updates) == 0 {
		return nil
	}

	return uc.repo.UpdateDemo(ctx, request.Id, updates)
}

func (uc *DemoUseCase) DeleteDemo(ctx context.Context, id string) error {
	return uc.repo.DeleteDemo(ctx, id)
}

func (uc *DemoUseCase) GetDemoCount(ctx context.Context, status *uint32) int64 {
	return uc.repo.GetDemoCount(ctx, status)
}
