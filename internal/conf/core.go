package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfigUtils,
	NewBootstrapConfig,
	NewLoggerConfig,

	NewConsulConfig,
	NewRedisConfig,
	NewMysqlConfig,
)
