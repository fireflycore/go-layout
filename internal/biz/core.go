package biz

import (
	"github.com/google/wire"
)

// ProviderSet 组装 biz 层 provider。
var ProviderSet = wire.NewSet(
	NewDemoUseCase,
)
