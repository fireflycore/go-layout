package main

import (
	"context"
	"errors"
	"go-layout/internal/conf"
	"go-layout/internal/server"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// App 聚合 go-layout 模板服务的托管运行依赖。
type App struct {
	BootstrapConf *conf.BootstrapConf

	ManagementServer *server.ManagementServer
	ManagedServer    *agent.ManagedServer
	GrpcServer       *server.GrpcServer

	Lifecycle *agent.ServiceLifecycle
	Services  []*grpc.ServiceDesc

	Logger *logger.ServerLogger
}

// Run 使用统一托管入口运行服务，并在退出时自动完成 sidecar 收尾。
func (ist *App) Run(ctx context.Context) error {
	// 启动前先输出服务基础元信息，便于排查部署实例与监听地址。
	ist.Logger.Info("starting go-layout service",
		zap.String("service_name", ist.BootstrapConf.GetServiceName()),
		zap.String("service_namespace", ist.BootstrapConf.GetServiceNamespace()),
		zap.String("service_instance_id", ist.BootstrapConf.GetServiceInstanceId()),
		zap.String("grpc_addr", ist.GrpcServer.Addr()),
		zap.String("management_addr", ist.ManagementServer.Addr()),
		zap.String("otel_endpoint", ist.BootstrapConf.GetOtelEndpoint()),
		zap.Bool("otel_traces", ist.BootstrapConf.GetOtelTraces()),
		zap.Bool("otel_metrics", ist.BootstrapConf.GetOtelMetrics()),
		zap.Bool("otel_logs", ist.BootstrapConf.GetOtelLogs()),
	)
	if ist.Lifecycle != nil {
		// 在真正启动前先输出一份当前 sidecar 状态，便于确认生命周期对象已装配完成。
		status := ist.Lifecycle.Status()
		ist.Logger.Info("sidecar lifecycle prepared",
			zap.String("service_name", ist.BootstrapConf.GetServiceName()),
			zap.String("service_instance_id", ist.BootstrapConf.GetServiceInstanceId()),
			zap.Bool("connected", status.Connected),
			zap.Bool("registered", status.Registered),
		)
	}
	// 输出启动完成日志，说明管理端口和业务端口已准备好进入托管运行阶段。
	ist.Logger.Info("go-layout service startup completed",
		zap.String("grpc_addr", ist.GrpcServer.Addr()),
		zap.String("management_addr", ist.ManagementServer.Addr()),
		zap.String("service_endpoint", ist.BootstrapConf.GetServiceEndpoint()),
	)
	if ist.ManagedServer == nil {
		// 没有统一托管器时无法保证 sidecar 生命周期和本地服务协同退出。
		return errors.New("managed server is required")
	}
	// 把主运行控制权交给 ManagedServer，统一托管 gRPC、管理端口和 sidecar 生命周期。
	err := ist.ManagedServer.Run(ctx)
	// 托管运行结束后输出最终停止日志。
	ist.Logger.Info("go-layout service stopped",
		zap.String("service_name", ist.BootstrapConf.GetServiceName()),
		zap.String("service_instance_id", ist.BootstrapConf.GetServiceInstanceId()),
	)
	return err
}

// NewApp 组装应用根对象，把业务端口与 sidecar 托管器集中到同一个启动入口。
func NewApp(bootstrapConf *conf.BootstrapConf, managementServer *server.ManagementServer, managedServer *agent.ManagedServer, grpcServer *server.GrpcServer, lifecycle *agent.ServiceLifecycle, services []*grpc.ServiceDesc, logger *logger.ServerLogger) *App {
	return &App{
		BootstrapConf: bootstrapConf,

		ManagementServer: managementServer,
		ManagedServer:    managedServer,
		GrpcServer:       grpcServer,

		Lifecycle: lifecycle,
		Services:  services,

		Logger: logger,
	}
}
