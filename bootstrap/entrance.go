package bootstrap

import (
	_ "embed"
	"fmt"
	"github.com/lhdhtrc/microservice-go/db"
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro/etcd"
	"github.com/lhdhtrc/microservice-go/micro/grpc"
	"github.com/lhdhtrc/microservice-go/remote"
	"github.com/lhdhtrc/microservice-go/utils/process"
	"microservice-go/plugin"
	"microservice-go/register"
	"microservice-go/store"
	"runtime"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config = plugin.SetupViper(&CONFIG)
	store.Use.Logger = logger.New(&store.Use.Config.Logger)

	store.Use.Grpc = grpc.New(store.Use.Logger)
	store.Use.Remote = remote.New(store.Use.Logger)

	dbs := db.New(store.Use.Logger)

	/********************************* read remote config ---- start *********************************/
	var etcdConfig db.ConfigEntity
	remoteConfig := []string{}
	store.Use.Remote.ReadRemoteConfig(remoteConfig, []interface{}{
		&etcdConfig,
	})
	/********************************* read remote config ---- end *********************************/

	/********************************* get remote cert ---- start *********************************/
	store.Use.Remote.GetRemoteCert("etcd", &etcdConfig.Tls)
	/********************************* get remote cert ---- end *********************************/

	/********************************* use etcd as microservice register ---- start *********************************/
	etcdCli := dbs.SetupEtcd(&etcdConfig)
	store.Use.Micro = etcd.New(etcdCli, store.Use.Logger, &store.Use.Config.Micro)
	/********************************* use etcd as microservice register ---- end *********************************/

	/********************************* service retry ---- start *********************************/
	store.Use.Micro.WithRetryBefore(func() {
		store.Use.Logger.Remote = nil
	})
	store.Use.Micro.WithRetryAfter(func() {
		store.Use.Grpc.Server.Stop()
		register.ServiceInstance()
	})
	/********************************* service retry ---- start *********************************/

	/********************************* register service ---- start *********************************/
	store.Use.Micro.CreateLease()
	register.ServiceInstance()
	/********************************* register service ---- end *********************************/

	store.Use.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		store.Use.Micro.Deregister()
	})
}
