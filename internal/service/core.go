package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewConfigCenterRemoteService,
	NewAccessLoggerRemoteService,
	NewServerLoggerRemoteService,
	NewOperationLoggerRemoteService,

	NewDemoService,
)
