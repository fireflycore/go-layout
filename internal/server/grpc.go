package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/dep"
	"go-layout/internal/service"

	"github.com/fireflycore/go-micro/middleware"
	ggm "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func NewGrpcServer(
	logger dep.AccessLogger,

	demoService *service.DemoService,
) *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(ggm.ChainUnaryServer(
		middleware.PropagateIncomingMetadata,
		middleware.GrpcAccessLogger(logger),
	)))

	demo.RegisterDemoServiceServer(srv, demoService)

	return srv
}
