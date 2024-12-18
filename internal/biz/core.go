package biz

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewDemoUseCase)

type AccessLogger func(b []byte, msg string)
type ServerLogger func(b []byte)
type OperationLogger func(b []byte)

type OuterGrpcClient grpc.ClientConnInterface
type InternalGrpcClient grpc.ClientConnInterface
