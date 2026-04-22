package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfUtils,
	NewBootstrapConf,
	NewBootstrapConfImpl,

	NewConsulConf,
	NewRedisConf,
	NewMysqlConf,
)
