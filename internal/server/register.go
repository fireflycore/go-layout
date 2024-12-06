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

func NewRegisterServer(bc *conf.BootstrapConf, d *data.Data) (micro.Register, func(), error) {
	register, err := etcd.NewRegister(d.Etcd, bc.Micro)
	return register, func() {
		register.Uninstall()
		fmt.Println("uninstall all service for this node from the register")
	}, err
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
