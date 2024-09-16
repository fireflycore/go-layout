package store

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
	"go-layout/config"
	"go-layout/model"
	"go.uber.org/zap"
)

type _CoreEntity struct {
	Config          *config.CoreEntity
	RemoteService   *model.RemoteServiceEntity
	ServiceDiscover map[string][]string

	Micro *micro.CoreEntity
	Etcd  *etcd.CoreEntity
	Task  *task.CoreEntity

	Logger *zap.Logger
}

var Use = new(_CoreEntity)
