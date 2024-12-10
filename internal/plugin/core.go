package plugin

import (
	"github.com/google/wire"
	logger "github.com/lhdhtrc/logger-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProviderSet = wire.NewSet(
	NewTask,
	NewLogger,

	NewGrpcClient,

	NewAccessLogger,
	NewServerLogger,
	NewOperationLogger,
)

func NewTask(bc *conf.BootstrapConf) *task.Instance {
	return task.New(bc.Task)
}

func NewLogger(bc *conf.BootstrapConf, handle biz.ServerLogger) *zap.Logger {
	return logger.New(bc.Logger, handle)
}

func NewGrpcClient(bc *conf.BootstrapConf) (*grpc.ClientConn, error) {
	var addr string
	if bc.Gateway.Network == bc.Micro.Network {
		addr = bc.Gateway.InternalNetAddr
	} else {
		addr = bc.Gateway.OuterNetAddr
	}
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
