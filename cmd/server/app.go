package main

import (
	"context"
	"errors"
	"go-layout/internal/conf"
	"go-layout/internal/server"
	"net"
	"strconv"

	"github.com/fireflycore/go-micro/logger"
	"go.uber.org/zap"
)

// App 聚合服务进程的托管运行依赖。
type App struct {
	BootstrapConfig *conf.BootstrapConfig

	Server *server.AppServer
	Logger *logger.ServerLogger
}

// Run 使用 SidecarAgent.Run 统一托管 gRPC、管理端口与 sidecar watch/replay，退出时自动收尾。
func (ist *App) Run(ctx context.Context) error {
	// 启动前先输出服务基础元信息，便于排查部署实例与监听地址。
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
		// 在真正启动前先输出一份当前 sidecar 状态，便于确认生命周期对象已装配完成。
		status := ist.Server.SidecarAgent.Status()
		ist.Logger.Info("sidecar agent prepared",
			zap.String("service_name", ist.BootstrapConfig.Service.Name),
			zap.String("service_instance_id", ist.BootstrapConfig.App.InstanceId),
			zap.Bool("connected", status.Connected),
			zap.Bool("registered", status.Registered),
		)
	}
	// 输出启动完成日志，说明管理端口和业务端口已准备好进入托管运行阶段。
	ist.Logger.Info("service startup completed",
		zap.String("grpc_addr", ist.Server.GrpcServer.Addr()),
		zap.String("management_addr", ist.Server.ManagedServer.Addr()),
		zap.String("service_endpoint", net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(ist.BootstrapConfig.ServerPort), 10))),
	)
	if ist.Server.SidecarAgent == nil {
		// 没有统一托管器时无法保证 sidecar 生命周期和本地服务协同退出。
		return errors.New("sidecar agent is required")
	}

	// 把主运行控制权交给 SidecarAgent，统一托管 gRPC、管理端口和 sidecar 生命周期。
	err := ist.Server.SidecarAgent.Run(ctx)

	// 托管运行结束后输出最终停止日志。
	ist.Logger.Info("service stopped",
		zap.String("service_name", ist.BootstrapConfig.Service.Name),
		zap.String("service_instance_id", ist.BootstrapConfig.App.InstanceId),
	)
	return err
}

// NewApp 组装应用根对象。
func NewApp(
	bootstrapConfig *conf.BootstrapConfig,
	server *server.AppServer,
	logger *logger.ServerLogger,
) *App {
	return &App{
		BootstrapConfig: bootstrapConfig,
		Server:          server,
		Logger:          logger,
	}
}
