package biz

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewDemoUseCase)

type AccessLogger func(b []byte, msg string)
type ServerLogger func(b []byte)
type OperationLogger func(b []byte)
