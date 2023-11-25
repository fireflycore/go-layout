package bootstrap

import (
	"context"
	_ "embed"
	"microservice-go/core"
	"microservice-go/register"
	"microservice-go/store"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config.Local = core.Use.SetupViper(&CONFIG)
	store.Use.Micro.Cli = core.Use.DB.SetupEtcd(&store.Use.Config.Local.Micro.RegisterCenter)
	store.Use.Micro.Ctx, store.Use.Micro.Cancel = context.WithCancel(context.Background())
	store.Use.Micro.Lease = core.Use.Micro.Etcd.CreateLease(store.Use.Micro.Cli)

	register.ServiceInstance()
}
