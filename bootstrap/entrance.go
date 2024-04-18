package bootstrap

import (
	_ "embed"
	"fmt"
	etcdCore "github.com/lhdhtrc/etcd-go/core"
	etcdModel "github.com/lhdhtrc/etcd-go/model"
	"github.com/lhdhtrc/func-go/process"
	loggerCore "github.com/lhdhtrc/logger-go/core"
	microEtcdCore "github.com/lhdhtrc/microcore-go/etcd"
	"github.com/lhdhtrc/microcore-go/grpc"
	taskCore "github.com/lhdhtrc/task-go/core"
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

	/********************************* setup logger core ---- start *********************************/
	store.Use.Config.Logger.AppId = store.Use.Config.System.AppId
	store.Use.Logger = loggerCore.Setup(store.Use.Config.Logger)
	/********************************* setup logger core ---- start *********************************/

	/********************************* setup task core ---- start *********************************/
	store.Use.Task = taskCore.New(store.Use.Config.Task)
	store.Use.Task.Setup()
	/********************************* setup task core ---- end *********************************/

	/********************************* create task ---- start *********************************/
	var etcdConfig etcdModel.ConfigEntity
	remoteConfig := []string{}
	task.ReadRemoteConfig(remoteConfig, []interface{}{
		&etcdConfig,
	})
	task.GetRemoteCert("etcd", &etcdConfig.Tls)
	/********************************* create task ---- end *********************************/

	/********************************* setup micro core ---- start *********************************/
	etcdCli := etcdCore.Setup(store.Use.Logger, &etcdConfig)
	store.Use.Grpc = grpc.New(store.Use.Logger)
	store.Use.Micro = microEtcdCore.New(etcdCli, store.Use.Logger, &store.Use.Config.Micro)
	store.Use.Micro.WithRetryAfter(func() {
		store.Use.Grpc.Server.Stop()
		api.ServiceInstance()
	})
	store.Use.Micro.CreateLease()
	api.ServiceInstance()
	/********************************* setup micro core ---- end *********************************/

	store.Use.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		store.Use.Micro.Deregister()
	})
}
