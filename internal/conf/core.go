package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfUtils,
	NewBootstrapConf,

	NewEtcdConfLoader,
	NewRedisConfLoader,
	NewMysqlConfLoader,

	NewEtcdConf,
	NewRedisConf,
	NewPostgresConf,
)
