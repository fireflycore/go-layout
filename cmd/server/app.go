package main

import (
	"context"
	"errors"
	"go-layout/internal/conf"
	"go-layout/internal/server"

	"github.com/fireflycore/go-micro/logger"
	"go.uber.org/zap"
)

// App 聚合服务进程的托管运行依赖。
type App struct {
	BootstrapConfig *conf.BootstrapConfig

	Server *server.AppServer
	Logger *logger.ServerLogger
}

// Run 使用 agent.Run 统一托管 gRPC、管理端口与 sidecar watch/replay，退出时自动收尾。
func (ist *App) Run(ctx context.Context) error {
	ist.Logger.Info("starting service",
		zap.String("service_name", ist.BootstrapConfig.Service.Name),
		zap.String("service_namespace", ist.BootstrapConfig.Service.Namespace),
		zap.String("service_instance_id", ist.BootstrapConfig.App.InstanceId),
		zap.String("grpc_addr", ist.Server.GrpcServer.Addr()),
		zap.String("management_addr", ist.Server.ManagedServer.Addr()),
		zap.String("otel_endpoint", ist.BootstrapConfig.Telemetry.OTLPEndpoint),
		zap.Bool("otel_traces", ist.BootstrapConfig.Telemetry.Traces),
		zap.Bool("otel_metrics", ist.BootstrapConfig.Telemetry.Metrics),
		zap.Bool("otel_logs", ist.BootstrapConfig.Telemetry.Logs),
	)
	if ist.Server.SidecarAgent != nil {
		status := ist.Server.SidecarAgent.Status()
		ist.Logger.Info("sidecar agent prepared",
			zap.String("service_name", ist.BootstrapConfig.Service.Name),
			zap.String("service_instance_id", ist.BootstrapConfig.App.InstanceId),
			zap.Bool("connected", status.Connected),
			zap.Bool("registered", status.Registered),
		)
	}
	ist.Logger.Info("service startup completed",
		zap.String("grpc_addr", ist.Server.GrpcServer.Addr()),
		zap.String("management_addr", ist.Server.ManagedServer.Addr()),
		zap.String("service_endpoint", ist.BootstrapConfig.ServiceEndpoint()),
	)
	if ist.Server.SidecarAgent == nil {
		return errors.New("sidecar agent is required")
	}

	err := ist.Server.SidecarAgent.Run(ctx)

	ist.Logger.Info("service stopped",
		zap.String("service_name", ist.BootstrapConfig.Service.Name),
		zap.String("service_instance_id", ist.BootstrapConfig.App.InstanceId),
	)
	return err
}

func NewApp(
	bootstrapConfig *conf.BootstrapConfig,
	appServer *server.AppServer,
	logger *logger.ServerLogger,
) *App {
	return &App{
		BootstrapConfig: bootstrapConfig,
		Server:          appServer,
		Logger:          logger,
	}
}
