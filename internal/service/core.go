package service

import (
	"github.com/google/wire"
)

// ProviderSet 组装 service 层 provider。
//
// 这一层只暴露 RPC 入口，不承载业务规则本身。
var ProviderSet = wire.NewSet(
	NewDemoService,
)
