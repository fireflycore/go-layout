package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/registry"
	"github.com/fireflycore/go-micro/registry/etcd"
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

func NewRegisterCenterRepo(register registry.Register, logger *zap.Logger) []*grpc.ServiceDesc {
	raw := []*grpc.ServiceDesc{
		&demo.DemoService_ServiceDesc,
	}

	if errs := registry.NewRegisterService(raw, register); len(errs) != 0 {
		logger.Error("register service failed", zap.Any("errors", errs))
	}

	return raw
}
