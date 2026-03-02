package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"

	etcd "github.com/fireflycore/go-etcd/registry"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewRegisterServer(bootstrapConf *conf.BootstrapConf, cli *clientv3.Client) (registry.Register, error) {
	meta := registry.Meta{
		Env:     bootstrapConf.Env,
		AppId:   bootstrapConf.AppId,
		Version: bootstrapConf.Version,
	}

	return etcd.NewRegister(cli, &meta, bootstrapConf.Micro)
}

func NewRegisterCenterRepo(register registry.Register, logger *logger.Core) []*grpc.ServiceDesc {
	raw := []*grpc.ServiceDesc{
		&demo.DemoService_ServiceDesc,
	}

	if errs := registry.NewRegisterService(raw, register); len(errs) != 0 {
		logger.Error("register service failed", zap.Errors("errors", errs))
	}

	return raw
}
