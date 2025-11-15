package biz

import (
	"go-layout/internal/biz/repo"
	"go-layout/internal/dto/convert"
)

type DemoUseCase struct {
	repo repo.DemoRepo

	dto convert.DemoConverter
}

func NewDemoUseCase(repo repo.DemoRepo, dto convert.DemoConverter) *DemoUseCase {
	return &DemoUseCase{
		repo: repo,

		dto: dto,
	}
}
