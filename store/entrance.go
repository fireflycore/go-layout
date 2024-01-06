package store

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
)

type _Entrance struct {
	Config  *config.EntranceEntity
	Logger  logger.Abstraction
	Micro   micro.Abstraction
	Grpc    *grpc.EntranceEntity
	Service map[string][]string

	Remote *remote.EntranceEntity
}

var Use = new(_Entrance)
