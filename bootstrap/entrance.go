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
	"microservice-go/service"
	"microservice-go/store"
)

//go:embed file/config.yaml
var CONFIG []byte

func Setup() {
	store.Use.Config = plugin.SetupViper(&CONFIG)
	store.Use.Logger = logger.New(&store.Use.Config.Logger)

	store.Use.Grpc = grpc.New(store.Use.Logger)
	store.Use.Remote = remote.New(store.Use.Logger)

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

	dbs := db.New(store.Use.Logger)

	/********************************* use etcd as microservice register ---- start *********************************/
	etcdCli := dbs.SetupEtcd(&etcdConfig)
	store.Use.Micro = etcd.New(etcdCli, store.Use.Logger, &store.Use.Config.Micro)
	/********************************* use etcd as microservice register ---- end *********************************/

	/********************************* discover service ---- start *********************************/
	store.Use.Service = make(map[string][]string)
	store.Use.Micro.Watcher(&[]string{
		"/microservice/lhdht/logger/Add",
	}, &store.Use.Service)

	fmt.Println(store.Use.Service)
	store.Use.Logger.Remote = service.Use.Logger.Add
	/********************************* discover service ---- end *********************************/

	/********************************* register service ---- start *********************************/
	store.Use.Micro.CreateLease()
	register.ServiceInstance()
	/********************************* register service ---- end *********************************/

	process.Watcher(func() {
		store.Use.Micro.Deregister()
	})
}
