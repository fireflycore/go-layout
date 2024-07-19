package store

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
	"go.uber.org/zap"
	"microservice-go/config"
	"microservice-go/model"
)

type _CoreEntity struct {
	ServiceDiscover map[string][]string
	RemoteService   *model.RemoteServiceEntity

	Config *config.CoreEntity
	Logger *zap.Logger
	Micro  *micro.CoreEntity
	Etcd   *etcd.CoreEntity
	Task   *task.CoreEntity
}

var Use = new(_CoreEntity)
