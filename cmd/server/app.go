package main

import (
	"fmt"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
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
	register.WithRetryBefore(func() {
		fmt.Println("重试之前的函数")
	})
	register.WithRetryAfter(func() {
		micro.NewRegisterService(services, register)
		go register.SustainLease()
		fmt.Println("重试之后的函数")
	})
	go register.SustainLease()

	return &App{
		Listener:   nl,
		GrpcServer: gs,
		Register:   register,
		Services:   services,
		Logger:     logger,
	}
}
