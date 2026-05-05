package dep

import (
	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/telemetry"
)

func NewTelemetryProviders(bootstrapConfig *conf.BootstrapConfig) (*telemetry.Providers, error) {
	return telemetry.NewProviders(&bootstrapConfig.Telemetry, &telemetry.Resource{
		ServiceId:         bootstrapConfig.App.Id,
		ServiceName:       bootstrapConfig.Service.Name,
		ServiceVersion:    bootstrapConfig.App.Version,
		ServiceNamespace:  bootstrapConfig.Service.Namespace,
		ServiceInstanceId: bootstrapConfig.App.InstanceId,
	})
}
