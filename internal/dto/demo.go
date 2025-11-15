package dto

import (
	"go-layout/depend/goverter/biz"
	"go-layout/internal/dto/convert"
)

func NewDemoConvert() convert.DemoConverter {
	return &biz.DemoConverterImpl{}
}
