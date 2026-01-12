package dep

import (
	"github.com/fireflycore/go-utils/compress"
	"github.com/fireflycore/go-utils/crypto"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRemoteServiceGrpcClient,

	NewLogger,
	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,

	crypto.NewAESCrypto,
	compress.NewGZIP,
)
