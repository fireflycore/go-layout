package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAccessLoggerRemoteService,
	NewServerLoggerRemoteService,
	NewOperationLoggerRemoteService,

	NewDemoService,
)
