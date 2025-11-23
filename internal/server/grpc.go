package server

import (
	"github.com/lhdhtrc/micro-go/pkg/middleware"
	demo "go-layout/depend/protobuf/gen/acme/demo/v1"
	"go-layout/internal/depend"
	"go-layout/internal/service"
	"google.golang.org/grpc"
)

func NewGrpcServer(
	logger depend.AccessLogger,

	Demo *service.DemoService,
) *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(ggm.ChainUnaryServer(
		middleware.PropagateIncomingMetadata,
		middleware.GrpcAccessLogger(logger),
	)))

	demo.RegisterDemoServiceServer(srv, Demo)

	return srv
}
