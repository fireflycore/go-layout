package server

import (
	"github.com/google/wire"
)

// ProviderSet 组装 server 层 provider。
var ProviderSet = wire.NewSet(
	NewGrpcServer,
	NewServiceDesc,

	NewSidecarAgent,
	NewSidecarStatusProvider,

	NewAppServer,
	NewAppManagedServer,
)
