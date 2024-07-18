package bootstrap

import (
	_ "embed"
	"fmt"
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	"github.com/lhdhtrc/func-go/process"
	loger "github.com/lhdhtrc/logger-go/pkg"
	"github.com/lhdhtrc/task-go/pkg"
	"microservice-go/api"
	"microservice-go/plugin"
	"microservice-go/store"
	"runtime"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config = plugin.InstallViper(&CONFIG)
	store.Use.RemoteService = plugin.InstallRemoteService(store.Use.Config.System.RemoteService)

	/********************************* logger core ---- start *********************************/
	store.Use.Logger = loger.New(&store.Use.Config.Logger, plugin.InstallRemoteLogger)
	/********************************* logger core ---- start *********************************/

	/********************************* task ---- start *********************************/
	store.Use.Task = task.New(&store.Use.Config.Task)

	var etcdConfig etcd.ConfigEntity
	store.Use.Task.InitConfig([]string{}, []interface{}{
		&etcdConfig,
	})
	store.Use.Task.Await()
	store.Use.Task.InitCert("etcd", &etcdConfig.Tls)
	store.Use.Task.Await()
	/********************************* task ---- end *********************************/

	/********************************* micro core ---- start *********************************/
	store.Use.Etcd = etcd.New(store.Use.Logger, &etcdConfig)
	store.Use.Etcd.WithLeaseRetryAfter(func() {
		store.Use.GrpcServer.Stop()
		store.Use.GrpcServer = GrpcServer(api.ServiceInstance, store.Use.Config.System.RunPort)
	})
	store.Use.Etcd.InitLease()
	store.Use.GrpcServer = GrpcServer(api.ServiceInstance, store.Use.Config.System.RunPort)
	/********************************* micro core ---- end *********************************/

	store.Use.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		store.Use.Etcd.Uninstall()
		store.Use.Task.Uninstall()
	})
}
