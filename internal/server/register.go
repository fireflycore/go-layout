package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewServiceDesc 收敛当前模板对外暴露的 gRPC 服务描述。
func NewServiceDesc(log *logger.ServerLogger) []*grpc.ServiceDesc {
	raw := []*grpc.ServiceDesc{
		&demo.DemoService_ServiceDesc,
	}

	log.Info("grpc service descriptors ready", zap.Int("count", len(raw)))
	return raw
}

// NewServiceLifecycle 把 gRPC 服务描述转换成 sidecar 生命周期桥接对象。
func NewServiceLifecycle(bootstrapConf *conf.BootstrapConf, services []*grpc.ServiceDesc) (*agent.ServiceLifecycle, error) {
	return agent.NewServiceLifecycleFromGRPC(agent.GRPCDescriptorOptions{
		AppId:       bootstrapConf.AppId,
		AppName:     bootstrapConf.AppName,
		ServiceName: bootstrapConf.GetServiceName(),
		Namespace:   bootstrapConf.GetServiceNamespace(),
		DNS:         bootstrapConf.GetServiceDNS(),
		Env:         bootstrapConf.Env,
		Port:        int(bootstrapConf.GetServerPort()),
		Protocol:    "grpc",
		Version:     bootstrapConf.Version,
		ServiceOptions: &agent.ServiceOptions{
			InstanceId: bootstrapConf.GetServiceInstanceId(),
			Namespace:  bootstrapConf.GetServiceNamespace(),
		},
		RawServices: services,
	}, agent.DefaultLocalRuntimeOptions(bootstrapConf.GetSidecarAgentBaseURL()), agent.LifecycleOptions{
		GracePeriod: bootstrapConf.GetSidecarGracePeriod(),
	})
}

// NewSidecarStatusProvider 复用生命周期对象本身作为 sidecar 状态提供者。
func NewSidecarStatusProvider(lifecycle *agent.ServiceLifecycle) SidecarStatusProvider {
	return lifecycle
}
