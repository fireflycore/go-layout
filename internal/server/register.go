package server

import (
	"fmt"
	micro "github.com/lhdhtrc/micro-go/pkg"
	"github.com/lhdhtrc/micro-go/pkg/etcd"
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"
	"go-layout/internal/data"
	"google.golang.org/grpc"
)

func NewRegisterServer(bc *conf.BootstrapConf, d *data.Data) (micro.Register, error) {
	return etcd.NewRegister(d.Etcd, bc.Micro)
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
