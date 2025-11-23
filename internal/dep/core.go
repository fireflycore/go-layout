package dep

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRemoteServiceGrpcClient,

	NewLogger,
	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,
)
