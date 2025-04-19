package dep

import (
	"github.com/google/wire"
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	"go-layout/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProviderSet = wire.NewSet(
	NewEtcd,

	NewRemoteServiceGrpcClient,

	NewLogger,
	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,
)

func NewEtcd(ec *etcd.Config) (*clientv3.Client, error) {
	return etcd.New(ec)
}

func NewRemoteServiceGrpcClient(bc *conf.BootstrapConf) (*grpc.ClientConn, error) {
	var addr string
	if bc.Gateway.Network.SN == bc.Micro.Network.SN {
		addr = bc.Gateway.Network.Internal
	} else {
		addr = bc.Gateway.Network.External
	}
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
