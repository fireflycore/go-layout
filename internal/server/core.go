package server

import (
	"github.com/google/wire"
)

// ProviderSet 组装 server 层 provider。
//
// 这一层负责服务端口、管理端口与 sidecar 生命周期装配。
var ProviderSet = wire.NewSet(
	NewGrpcServer,
	NewServiceDesc,

	NewSidecarAgent,
	NewSidecarStatusProvider,

	NewAppServer,
	NewAppManagedServer,
)
