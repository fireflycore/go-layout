package server

import (
	"context"
	"errors"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/invocation"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/telemetry"
	"go.uber.org/zap"
)

// AppServer 汇总服务运行期需要持有的主对象。
type AppServer struct {
	ConnectionManager *invocation.ConnectionManager
	ManagedServer     *AppManagedServer
	SidecarAgent      *agent.Agent
	GrpcServer        *GrpcServer
}

// NewAppServer 将 gRPC、管理端口关闭逻辑写入 sidecar Agent（ConfigureRun），并返回同一 Agent 作为统一运行入口。
func NewAppServer(
	bootstrapConfig *conf.BootstrapConfig,

	log *logger.ServerLogger,
	providers *telemetry.Providers,

	grpcServer *GrpcServer,
	sidecarAgent *agent.Agent,
	managedServer *AppManagedServer,

	connectionManager *invocation.ConnectionManager,
) (*AppServer, error) {
	if sidecarAgent == nil {
		return nil, errors.New("sidecar agent is required")
	}

	// 统一把 gRPC 与管理端口的启动、关闭动作挂到 sidecar 生命周期上。
	sidecarAgent.ConfigureRun(agent.SidecarAgentConfig{
		Serve: func(ctx context.Context) error {
			errCh := make(chan error, 2)
			// 管理端口与 gRPC 端口并行启动，任何一个异常退出都视为服务异常。
			go func() {
				errCh <- managedServer.Serve()
			}()
			go func() {
				errCh <- grpcServer.Serve()
			}()
			select {
			case err := <-errCh:
				// 这里保留服务名与实例号，便于定位是哪一个实例先退出。
				if err != nil {
					log.Error("service server exited with error",
						zap.String("service_name", bootstrapConfig.Service.Name),
						zap.String("service_instance_id", bootstrapConfig.App.InstanceId),
						zap.Error(err),
					)
				}
				return err
			case <-ctx.Done():
				return nil
			}
		},
		Shutdown: func(ctx context.Context) error {
			// 先停止对外监听，再回收出站连接和可观测性 provider。
			log.Info("stopping service servers",
				zap.String("service_name", bootstrapConfig.Service.Name),
				zap.String("service_instance_id", bootstrapConfig.App.InstanceId),
			)
			managedServer.Stop()
			grpcServer.Stop()
			if connectionManager != nil {
				if err := connectionManager.Close(); err != nil {
					log.Warn("failed to close invocation connection manager",
						zap.String("service_name", bootstrapConfig.Service.Name),
						zap.String("service_instance_id", bootstrapConfig.App.InstanceId),
						zap.Error(err),
					)
				}
			}
			providers.Shutdown()
			log.Info("service servers stopped",
				zap.String("service_name", bootstrapConfig.Service.Name),
				zap.String("service_instance_id", bootstrapConfig.App.InstanceId),
			)
			return nil
		},
	})

	return &AppServer{
		ConnectionManager: connectionManager,
		ManagedServer:     managedServer,
		SidecarAgent:      sidecarAgent,
		GrpcServer:        grpcServer,
	}, nil
}
