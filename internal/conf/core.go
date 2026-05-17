package conf

import "github.com/google/wire"

// ProviderSet 组装 conf 层 provider。
//
// 这一层只负责启动期静态配置解析，不承载运行时配置拉取。
var ProviderSet = wire.NewSet(
	NewConfigUtils,
	NewBootstrapConfig,
	NewLoggerConfig,

	NewConsulConfig,
	NewRedisConfig,
	NewMysqlConfig,
)
