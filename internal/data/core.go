package data

import (
	"github.com/google/wire"
)

// ProviderSet 组装 data 层 provider。
var ProviderSet = wire.NewSet(
	NewData,
	NewConsul,
	NewRedis,
	NewMysql,
	NewConfigStore,

	NewDemoRepo,
)
