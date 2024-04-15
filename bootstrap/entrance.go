package bootstrap

import (
	_ "embed"
	"fmt"
	"github.com/lhdhtrc/microservice-go/db"
	"github.com/lhdhtrc/microservice-go/micro/etcd"
	"github.com/lhdhtrc/microservice-go/micro/grpc"
	"github.com/lhdhtrc/microservice-go/utils/process"
	taskCore "github.com/lhdhtrc/task-go/core"
	taskModel "github.com/lhdhtrc/task-go/model"
	"microservice-go/api"
	"microservice-go/plugin"
	"microservice-go/store"
	"microservice-go/task"
	"runtime"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config = plugin.SetupViper(&CONFIG)

	store.Use.Task = taskCore.New(taskModel.ConfigEntity{
		MaxCache:       10,
		MaxConcurrency: 1,
		MinConcurrency: 0,
	})

	store.Use.Grpc = grpc.New(store.Use.Logger)

	dbs := db.New(store.Use.Logger)

	/********************************* read remote config ---- start *********************************/
	var etcdConfig db.ConfigEntity
	remoteConfig := []string{}
	task.ReadRemoteConfig(remoteConfig, []interface{}{
		&etcdConfig,
	})
	/********************************* read remote config ---- end *********************************/

	/********************************* get remote cert ---- start *********************************/
	task.GetRemoteCert("etcd", &etcdConfig.Tls)
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
		api.ServiceInstance()
	})
	/********************************* service retry ---- start *********************************/

	/********************************* register service ---- start *********************************/
	store.Use.Micro.CreateLease()
	api.ServiceInstance()
	/********************************* register service ---- end *********************************/

	store.Use.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		store.Use.Micro.Deregister()
	})
}
