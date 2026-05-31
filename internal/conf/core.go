package conf

import "github.com/google/wire"

// ProviderSet 组装 conf 层 provider。
var ProviderSet = wire.NewSet(
	NewConfigUtils,
	NewBootstrapConfig,

	NewLoggerConfig,

	NewConsulConfig,
	NewRedisConfig,
	NewMysqlConfig,
)
