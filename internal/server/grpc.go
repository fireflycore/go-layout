package server

import (
	"github.com/lhdhtrc/micro-go/pkg/middleware"
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz"
	"go-layout/internal/service"
	"google.golang.org/grpc"
)

func NewGrpcServer(
	logger biz.AccessLogger,

	Demo *service.DemoService,
) *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(
		middleware.GrpcAccessLogger(logger),
	))

	demo.RegisterDemoServiceServer(srv, Demo)

	return srv
}
