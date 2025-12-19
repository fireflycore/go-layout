package dep

import (
	"github.com/google/wire"
	compress "github.com/lhdhtrc/compress-go/pkg"
	crypto "github.com/lhdhtrc/crypto-go/pkg"
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
