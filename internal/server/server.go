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

// AppServer 把业务 gRPC、管理端口和 sidecar-agent 收敛成统一运行入口。
type AppServer struct {
	ConnectionManager *invocation.ConnectionManager
	ManagedServer     *AppManagedServer
	SidecarAgent      *agent.Agent
	GrpcServer        *GrpcServer
}

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

	sidecarAgent.ConfigureRun(agent.SidecarAgentConfig{
		GracePeriod: bootstrapConfig.SidecarAgentConfig().GracePeriod,
		Serve: func(ctx context.Context) error {
			errCh := make(chan error, 2)

			go func() {
				errCh <- managedServer.Serve()
			}()
			go func() {
				errCh <- grpcServer.Serve()
			}()

			select {
			case err := <-errCh:
				if err != nil {
					log.Error("managed server exited with error",
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
		Shutdown: func(context.Context) error {
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
			if providers != nil {
				_ = providers.Shutdown()
			}

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
