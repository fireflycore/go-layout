package bootstrap

import (
	"context"
	_ "embed"
	"microservice-go/core"
	"microservice-go/micro"
	"microservice-go/model/base"
	"microservice-go/register"
	"microservice-go/service"
	"microservice-go/store"
	"microservice-go/utils/logger"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config.Local = core.Use.SetupViper(&CONFIG)
	store.Use.Config.Local.Logger.WithRemote = service.Use.Logger.Add

	store.Use.Logger.Func = logger.New(store.Use.Config.Local.Logger)

	store.Use.Micro.Service = make(map[string][]string)
	store.Use.Micro.Cli = core.Use.DB.SetupEtcd(&store.Use.Config.Local.Micro.RegisterCenter)
	store.Use.Micro.Ctx, store.Use.Micro.Cancel = context.WithCancel(context.Background())
	store.Use.Micro.Lease = micro.Use.Etcd.CreateLease(store.Use.Micro.Cli)

	register.ServiceInstance()

	microserviceWatchConfig := []base.MicroserviceDiscoverEntity{
		{
			Gateway:   store.Use.Config.Local.Micro.GatewayName,
			Namespace: "logger",
		},
	}
	micro.Use.Etcd.Watcher(&microserviceWatchConfig)

	core.Use.WatchProcess(func() {
		micro.Use.Etcd.Deregister()
		store.Use.Micro.Cancel()
	})
}
