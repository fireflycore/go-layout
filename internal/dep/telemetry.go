package dep

import (
	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/telemetry"
)

// NewTelemetryProviders 创建服务级可观测性 provider 集合。
func NewTelemetryProviders(bootstrapConfig *conf.BootstrapConfig) (*telemetry.Providers, error) {
	return telemetry.NewProviders(&bootstrapConfig.Telemetry, &telemetry.Resource{
		ServiceId: bootstrapConfig.App.Id,
		// Environment 服务运行环境
		Environment: bootstrapConfig.App.Env,
		// ServiceName 服务名称
		ServiceName: bootstrapConfig.Service.Name,
		// ServiceVersion 服务版本
		ServiceVersion: bootstrapConfig.App.Version,
		// ServiceNamespace 服务命名空间
		ServiceNamespace: bootstrapConfig.Service.Namespace,
		// ServiceInstanceId 服务实例id
		ServiceInstanceId: bootstrapConfig.App.InstanceId,
	})
}
