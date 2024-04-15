package store

import (
	"github.com/lhdhtrc/microcore-go/grpc"
	microModel "github.com/lhdhtrc/microcore-go/model"
	"github.com/lhdhtrc/microservice-go/logger"

	taskCore "github.com/lhdhtrc/task-go/core"
	"go.uber.org/zap"
	"microservice-go/config"
)

type _Entrance struct {
	Config       *config.EntranceEntity
	Logger       *zap.Logger
	LoggerTemp   []logger.Entity
	Micro        microModel.Abstraction
	Grpc         *grpc.EntranceEntity
	GrpcEndpoint map[string][]string

	Task *taskCore.TaskCoreEntity
}

var Use = new(_Entrance)
