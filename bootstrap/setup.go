package bootstrap

import (
	_ "embed"
	"fmt"
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	"github.com/lhdhtrc/func-go/process"
	loger "github.com/lhdhtrc/logger-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg"
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

	/********************************* logger ---- start *********************************/
	store.Use.Logger = loger.New(&store.Use.Config.Logger, plugin.InstallServerLogger)
	/********************************* logger ---- start *********************************/

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

	/********************************* micro ---- start *********************************/
	store.Use.Micro = micro.New(store.Use.Logger)
	store.Use.Etcd = etcd.New(store.Use.Logger, &etcdConfig)
	store.Use.Etcd.WithLeaseRetryAfter(func() {
		store.Use.Micro.UninstallServer()
		store.Use.Micro.InstallServer(api.ServiceInstance, store.Use.Config.System.RunPort)
	})
	store.Use.Etcd.InitLease()
	store.Use.Micro.InstallServer(api.ServiceInstance, store.Use.Config.System.RunPort)
	/********************************* micro ---- end *********************************/

	store.Use.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		store.Use.Etcd.Uninstall()
		store.Use.Task.Uninstall()
	})
}
