package main

import (
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	"go-layout/internal/conf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	BootstrapConf *conf.BootstrapConf

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

func NewApp(bootstrapConf *conf.BootstrapConf, netListener net.Listener, grpcServer *grpc.Server, register micro.Register, services []*grpc.ServiceDesc, logger *zap.Logger) *App {
	register.WithRetryBefore(func() {
		logger.Info("retry before register")
	})
	register.WithRetryAfter(func() {
		micro.NewRegisterService(services, register)
		go register.SustainLease()
		logger.Info("retry after register")
	})
	go register.SustainLease()

	return &App{
		BootstrapConf: bootstrapConf,

		Listener:   netListener,
		GrpcServer: grpcServer,
		Register:   register,
		Services:   services,
		Logger:     logger,
	}
}
