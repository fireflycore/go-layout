package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"
	"go-layout/internal/dep"
	"go-layout/internal/service"

	gm "github.com/fireflycore/go-micro/middleware/grpc"
	ggm "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func NewGrpcServer(
	logger dep.AccessLogger,
	bootstrapConf *conf.BootstrapConf,

	demoService *service.DemoService,
) *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(ggm.ChainUnaryServer(
		recovery.UnaryServerInterceptor(),
		gm.NewServiceAccessLogger(logger),
		gm.NewInjectServiceContext(bootstrapConf),
	)))

	demo.RegisterDemoServiceServer(srv, demoService)

	return srv
}
