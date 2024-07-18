package store

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"microservice-go/config"
	"microservice-go/model"
)

type _StoreEntity struct {
	Config          *config.EntranceEntity
	Logger          *zap.Logger
	Etcd            *etcd.CoreEntity
	GrpcServer      *grpc.Server
	ServiceDiscover map[string][]string
	RemoteService   *model.RemoteServiceEntity

	Task *task.CoreEntity
}

var Use = new(_StoreEntity)
