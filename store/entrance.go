package store

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
	"github.com/lhdhtrc/microservice-go/micro/grpc"
	"github.com/lhdhtrc/microservice-go/remote"
	"microservice-go/config"
)

type _Entrance struct {
	Config     *config.EntranceEntity
	Logger     *logger.EntranceEntity
	LoggerTemp []logger.Entity
	Micro      micro.Abstraction
	Grpc       *grpc.EntranceEntity
	Endpoint    map[string][]string

	Remote *remote.EntranceEntity
}

var Use = new(_Entrance)
