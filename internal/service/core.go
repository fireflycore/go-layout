package service

import (
	"github.com/google/wire"
)

// ProviderSet 组装 service 层 provider。
var ProviderSet = wire.NewSet(
	NewDemoService,
)
