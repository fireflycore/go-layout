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
	microTask "github.com/lhdhtrc/microcore-go/task"
	taskCore "github.com/lhdhtrc/task-go/core"
	"microservice-go/api"
	"microservice-go/plugin"
	"microservice-go/store"
	"runtime"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config = plugin.SetupViper(&CONFIG)
	store.Use.RemoteService = plugin.SetupRemoteService(store.Use.Config.System.RemoteService)

	/********************************* setup logger core ---- start *********************************/
	store.Use.Logger = loggerCore.Setup(store.Use.Config.Logger, plugin.SetupRemoteLogger)
	/********************************* setup logger core ---- start *********************************/

	/********************************* setup task core ---- start *********************************/
	store.Use.Task = taskCore.New(store.Use.Config.Task)
	store.Use.Task.Setup()
	/********************************* setup task core ---- end *********************************/

	/********************************* create task ---- start *********************************/
	var etcdConfig etcdModel.ConfigEntity
	remoteConfig := []string{}
	microTask.ReadRemoteConfig(store.Use.Task, remoteConfig, []interface{}{
		&etcdConfig,
	})
	store.Use.Task.Await(1)
	microTask.GetRemoteCert(store.Use.Task, "etcd", &etcdConfig.Tls)
	store.Use.Task.Await(1)
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
