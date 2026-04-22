package server

import (
	"context"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/telemetry"
	"go.uber.org/zap"
)

// NewAppManagedServer 把业务 gRPC、管理端口和 sidecar 生命周期收敛成统一托管入口。
func NewAppManagedServer(
	bootstrapConf *conf.BootstrapConf,
	log *logger.ServerLogger,
	providers *telemetry.Providers,
	lifecycle *agent.ServiceLifecycle,
	managementServer *ManagementServer,
	grpcServer *GrpcServer,
) (*agent.ManagedServer, error) {
	return agent.NewManagedServer(agent.ManagedServerOptions{
		Lifecycle: lifecycle,
		Serve: func(ctx context.Context) error {
			errCh := make(chan error, 2)

			go func() {
				errCh <- managementServer.Serve()
			}()
			go func() {
				errCh <- grpcServer.Serve()
			}()

			select {
			case err := <-errCh:
				if err != nil {
					log.Error("managed server exited with error",
						zap.String("service_name", bootstrapConf.GetServiceName()),
						zap.String("service_instance_id", bootstrapConf.GetServiceInstanceId()),
						zap.Error(err),
					)
				}
				return err
			case <-ctx.Done():
				return nil
			}
		},
		Shutdown: func(ctx context.Context) error {
			log.Info("stopping service servers",
				zap.String("service_name", bootstrapConf.GetServiceName()),
				zap.String("service_instance_id", bootstrapConf.GetServiceInstanceId()),
			)

			managementServer.Stop()
			grpcServer.Stop()
			if providers != nil {
				_ = providers.Shutdown()
			}

			log.Info("service servers stopped",
				zap.String("service_name", bootstrapConf.GetServiceName()),
				zap.String("service_instance_id", bootstrapConf.GetServiceInstanceId()),
			)
			return nil
		},
	})
}
