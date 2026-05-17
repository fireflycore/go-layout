package dto

import (
	"github.com/google/wire"
)

// ProviderSet 组装 DTO 转换实现。
var ProviderSet = wire.NewSet(
	NewDemoDTO,
)
