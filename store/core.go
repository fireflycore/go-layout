package store

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg"
	tpg "github.com/lhdhtrc/task-go/pkg"
	"go.uber.org/zap"
	"microservice-go/config"
	"microservice-go/model"
)

type _CoreEntity struct {
	Config *config.CoreEntity
	Logger *zap.Logger
	Micro  *micro.CoreEntity
	Etcd   *etcd.CoreEntity

	ServiceDiscover map[string][]string
	RemoteService   *model.RemoteServiceEntity

	Task *tpg.CoreEntity
}

var Use = new(_CoreEntity)
