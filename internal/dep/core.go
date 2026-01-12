package dep

import (
	compress "github.com/fireflycore/go-utils/compress"
	crypto "github.com/fireflycore/go-utils/crypto"
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
