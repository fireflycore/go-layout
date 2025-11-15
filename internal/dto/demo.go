package dto

import (
	"go-layout/depend/dto"
	"go-layout/internal/dto/convert"
)

func NewDemoConvert() convert.DemoConverter {
	return &dto.DemoConverterImpl{}
}
