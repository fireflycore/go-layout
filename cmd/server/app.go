package main

import (
	"go-layout/internal/conf"
	"net"

	"github.com/fireflycore/go-micro/registry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	BootstrapConf *conf.BootstrapConf

	Listener   net.Listener
	GrpcServer *grpc.Server
	Register   registry.Register
	Services   []*grpc.ServiceDesc

	Logger *zap.Logger
}

func (ist *App) Start() {
	if err := ist.GrpcServer.Serve(ist.Listener); err != nil {
		panic(err)
	}
}

func (ist *App) Stop() {
	ist.Register.Uninstall()
	ist.GrpcServer.Stop()
}

func NewApp(bootstrapConf *conf.BootstrapConf, netListener net.Listener, grpcServer *grpc.Server, register registry.Register, services []*grpc.ServiceDesc, logger *zap.Logger) *App {
	register.WithRetryBefore(func() {
		logger.Info("retry before register")
	})
	register.WithRetryAfter(func() {
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
