package dep

import (
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-utils/compress"
	"github.com/fireflycore/go-utils/crypto"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRemoteServiceContext,
	NewRemoteServiceGrpcClient,

	NewLogger,
	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,

	sys.NewHostInfo,
	crypto.NewAESCrypto,
	compress.NewGZIP,
)
