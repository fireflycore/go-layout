package store

import (
	"github.com/lhdhtrc/microcore-go/grpc"
	microModel "github.com/lhdhtrc/microcore-go/model"
	taskCore "github.com/lhdhtrc/task-go/core"
	"go.uber.org/zap"
	"microservice-go/config"
	"microservice-go/model"
)

type _Entrance struct {
	Config        *config.EntranceEntity
	Logger        *zap.Logger
	Micro         microModel.MicroCoreInterface
	Grpc          *grpc.CoreEntity
	GrpcEndpoint  map[string][]string
	RemoteService *model.RemoteServiceEntity

	Task *taskCore.TaskCoreEntity
}

var Use = new(_Entrance)
