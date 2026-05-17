package biz

import (
	"github.com/google/wire"
)

// ProviderSet 组装 biz 层 provider。
//
// biz 层只承载业务规则、状态判断和跨 Repo 协调。
var ProviderSet = wire.NewSet(
	NewDemoUseCase,
)
