package main

import (
	micro "github.com/lhdhtrc/micro-go/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	Listener   net.Listener
	GrpcServer *grpc.Server
	Register   micro.Register
	Services   []*grpc.ServiceDesc

	Logger *zap.Logger
}

func (ist *App) Start() {
	if err := ist.GrpcServer.Serve(ist.Listener); err != nil {
		panic(err)
	}
}

func (ist *App) Stop() {
	ist.GrpcServer.Stop()
}

func NewApp(nl net.Listener, gs *grpc.Server, register micro.Register, services []*grpc.ServiceDesc, logger *zap.Logger) *App {
	return &App{
		Listener:   nl,
		GrpcServer: gs,
		Register:   register,
		Services:   services,
		Logger:     logger,
	}
}
