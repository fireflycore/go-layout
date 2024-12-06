package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/service"
	"google.golang.org/grpc"
)

func NewGrpcServer(
	Demo *service.DemoService,
) *grpc.Server {
	srv := grpc.NewServer()

	demo.RegisterDemoServiceServer(srv, Demo)

	return srv
}
