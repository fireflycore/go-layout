package dep

import (
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-micro/telemetry"
	"github.com/fireflycore/go-utils/compress"
	"github.com/fireflycore/go-utils/crypto"
	"github.com/google/wire"
)

// ProviderSet 组装模板库基础依赖，包括 telemetry、标准日志和 invocation 调用能力。
var ProviderSet = wire.NewSet(
	crypto.NewAESCrypto,
	compress.NewGZIP,
	sys.NewHostInfo,

	telemetry.NewProviders,

	logger.NewZapLogger,
	logger.NewAccessLogger,
	logger.NewServerLogger,

	NewInvocationConnectionManager,
	NewUnaryInvoker,
)
