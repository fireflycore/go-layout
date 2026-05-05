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

// NewSidecarAgent 基于 bootstrap 与 gRPC ServiceDesc 组装 go-consul/agent 主对象。
func NewSidecarAgent(bootstrapConfig *conf.BootstrapConfig, services []*grpc.ServiceDesc) (*agent.Agent, error) {
	serviceOpts := &agent.ServiceOptions{
		App:         bootstrapConfig.App,
		Kernel:      bootstrapConfig.Kernel,
		Service:     bootstrapConfig.Service,
		Protocol:    "grpc",
		ServerPort:  bootstrapConfig.ServerPort,
		ManagedPort: bootstrapConfig.ManagedPort,
	}

	sidecarConfig := bootstrapConfig.SidecarAgentConfig()
	sidecarConfig.RawServices = services

	return agent.New(serviceOpts, sidecarConfig)
}

func NewSidecarStatusProvider(a *agent.Agent) SidecarStatusProvider {
	return a
}
