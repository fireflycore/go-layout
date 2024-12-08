package plugin

import (
	"github.com/google/wire"
	task "github.com/lhdhtrc/task-go/pkg"
	"go-layout/internal/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProviderSet = wire.NewSet(
	NewTask,

	NewGrpcClient,

	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,
)

func NewTask(bc *conf.BootstrapConf) *task.Instance {
	return task.New(bc.Task)
}

func NewGrpcClient(bc *conf.BootstrapConf) (*grpc.ClientConn, error) {
	var addr string
	if bc.Gateway.Network == bc.Micro.Network {
		addr = bc.Gateway.InsideAddr
	} else {
		addr = bc.Gateway.OuterAddr
	}
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
