package dto

import (
	"go-layout/dep/dto"
	"go-layout/internal/biz/convert"
)

func NewDemoDTO() convert.DemoConvert {
	return &dto.DemoConvertImpl{}
}
