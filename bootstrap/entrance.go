package bootstrap

import (
	"context"
	_ "embed"
	"fmt"
	"microservice-go/core"
	"microservice-go/micro"
	"microservice-go/store"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config.Local = core.Use.SetupViper(&CONFIG)
	store.Use.Micro.Cli = core.Use.DB.SetupEtcd(&store.Use.Config.Local.Micro.RegisterCenter)
	store.Use.Micro.Ctx, store.Use.Micro.Cancel = context.WithCancel(context.Background())
	store.Use.Micro.Lease = core.Use.Micro.Etcd.CreateLease(store.Use.Micro.Cli)

	micro.RegisterServiceInstance()

	core.Use.WatchProcess(func() {
		if _, err := store.Use.Micro.Cli.Revoke(store.Use.Micro.Ctx, store.Use.Micro.Lease); err != nil {
			fmt.Println(err)
			return
		}
		store.Use.Micro.Cancel()
	})
}
