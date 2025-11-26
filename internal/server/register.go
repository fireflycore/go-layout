package server

import (
	"fmt"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	"github.com/lhdhtrc/micro-go/pkg/etcd"
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

func NewRegisterServer(bootstrapConf *conf.BootstrapConf, cli *clientv3.Client) (micro.Register, error) {
	meta := micro.Meta{
		Env:     bootstrapConf.Env,
		AppId:   bootstrapConf.AppId,
		Version: bootstrapConf.Version,
	}

	return etcd.NewRegister(cli, &meta, bootstrapConf.Micro)
}

func NewRegisterCenterRepo(register micro.Register) []*grpc.ServiceDesc {
	raw := []*grpc.ServiceDesc{
		&demo.DemoService_ServiceDesc,
	}

	if errs := micro.NewRegisterService(raw, register); len(errs) != 0 {
		fmt.Println(errs)
	}

	return raw
}
