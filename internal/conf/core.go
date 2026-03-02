package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfUtils,
	NewBootstrapConf,
	NewBootstrapConfImpl,

	NewEtcdConfLoader,
	NewRedisConfLoader,
	NewMysqlConfLoader,

	NewEtcdConf,
	NewRedisConf,
	NewMysqlConf,
)
