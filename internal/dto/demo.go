package dto

import (
	"go-layout/dep/dto"
	"go-layout/internal/biz/convert"
)

// NewDemoDTO 返回 Demo DTO 转换实现。
func NewDemoDTO() convert.DemoConvert {
	return &dto.DemoConvertImpl{}
}
