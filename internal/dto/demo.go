package dto

import (
	"go-layout/dep/goverter/biz"
	"go-layout/internal/dto/convert"
)

func NewDemoConvert() convert.DemoConverter {
	return &biz.DemoConverterImpl{}
}
