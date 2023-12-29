package bootstrap

import (
	_ "embed"
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro/etcd"
	"microservice-go/plugin"
	"microservice-go/store"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	config := plugin.SetupViper(&CONFIG)

	store.Use.Logger = logger.New(&logger.EntranceEntity{
		Config: config.Logger,
		Remote: nil,
	})
	store.Use.Micro = etcd.New(&etcd.EntranceEntity{
		Config:     nil,
		RetryCount: 0,
		Service:    nil,
		Ctx:        nil,
		Cancel:     nil,
		Cli:        nil,
		Lease:      0,
		Logger:     nil,
	})

}
