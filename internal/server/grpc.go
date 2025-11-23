package server

import (
	ggm "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/lhdhtrc/micro-go/pkg/middleware"
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/dep"
	"go-layout/internal/service"
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
