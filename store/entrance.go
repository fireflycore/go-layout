package store

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
	"github.com/lhdhtrc/microservice-go/micro/grpc"
	taskCore "github.com/lhdhtrc/task-go/core"
	"microservice-go/config"
)

type _Entrance struct {
	Config       *config.EntranceEntity
	Logger       *logger.EntranceEntity
	LoggerTemp   []logger.Entity
	Micro        micro.Abstraction
	Grpc         *grpc.EntranceEntity
	GrpcEndpoint map[string][]string

	Task *taskCore.TaskCoreEntity
}

var Use = new(_Entrance)
