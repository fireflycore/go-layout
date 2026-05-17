package data

import (
	"github.com/google/wire"
)

// ProviderSet 组装 data 层 provider。
//
// data 层负责数据库、缓存、Store 与 Repo 的具体实现。
var ProviderSet = wire.NewSet(
	NewData,
	NewConsul,
	NewRedis,
	NewMysql,
	NewConfigStore,

	NewDemoRepo,
)
