package dep

import (
	"github.com/fireflycore/go-micro/rpc"
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-utils/compress"
	"github.com/fireflycore/go-utils/crypto"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	crypto.NewAESCrypto,
	compress.NewGZIP,
	sys.NewHostInfo,

	rpc.NewRemoteInvokeServiceContext,
	rpc.NewRemoteServiceGrpcClient,

	NewLogger,
	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,
)
