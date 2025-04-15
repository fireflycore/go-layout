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

func NewRegisterServer(bc *conf.BootstrapConf, cli *clientv3.Client) (micro.Register, error) {
	return etcd.NewRegister(bc.AppId, cli, bc.Micro)
}

func NewRegisterCenterRepo(mr micro.Register) []*grpc.ServiceDesc {
	raw := []*grpc.ServiceDesc{
		&demo.DemoService_ServiceDesc,
	}

	if errs := micro.NewRegisterService(raw, mr); len(errs) != 0 {
		fmt.Println(errs)
	}

	return raw
}
