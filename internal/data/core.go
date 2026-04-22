package data

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewConsul,
	NewRedis,
	NewMysql,
	NewConfigStore,

	NewConfigRepo,

	NewDemoRepo,
)
