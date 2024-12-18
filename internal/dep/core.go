package dep

import (
	"github.com/google/wire"
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProviderSet = wire.NewSet(
	NewTask,
	NewEtcd,

	NewOuterGrpcClient,
	NewInternalGrpcClient,

	NewLogger,
	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,
)

func NewTask(bc *conf.BootstrapConf) *task.Instance {
	return task.New(bc.Task)
}

func NewEtcd(ec *etcd.Config) (*clientv3.Client, error) {
	return etcd.New(ec)
}

func NewOuterGrpcClient(bc *conf.BootstrapConf) (biz.OuterGrpcClient, error) {
	var addr string
	if bc.Gateway.Network == bc.Micro.Network {
		addr = bc.Gateway.InternalNetAddr
	} else {
		addr = bc.Gateway.OuterNetAddr
	}
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewInternalGrpcClient(bc *conf.BootstrapConf) (biz.InternalGrpcClient, error) {
	return grpc.NewClient(bc.Micro.InternalNetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
