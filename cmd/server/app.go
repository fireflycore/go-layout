package main

import (
	micro "github.com/lhdhtrc/micro-go/pkg"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	Listener   net.Listener
	GrpcServer *grpc.Server
	Register   micro.Register
	Services   []*grpc.ServiceDesc
}

func (ist *App) Start() {
	if err := ist.GrpcServer.Serve(ist.Listener); err != nil {
		panic(err)
	}
}

func (ist *App) Stop() {
	ist.GrpcServer.Stop()
}

func NewApp(nl net.Listener, gs *grpc.Server, register micro.Register, services []*grpc.ServiceDesc) *App {
	return &App{
		Listener:   nl,
		GrpcServer: gs,
		Register:   register,
		Services:   services,
	}
}
