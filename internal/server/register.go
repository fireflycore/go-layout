package server

import (
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewServiceDesc 汇总当前进程需要注册到 sidecar 的 gRPC 服务描述。
func NewServiceDesc(log *logger.ServerLogger) []*grpc.ServiceDesc {
	raw := []*grpc.ServiceDesc{
		&demo.DemoService_ServiceDesc,
	}

	// 启动期打印服务描述数量，方便排查 sidecar 注册是否漏服务。
	log.Info("grpc service descriptors ready", zap.Int("count", len(raw)))
	return raw
}

// NewSidecarAgent 基于 bootstrap 与 gRPC ServiceDesc 组装 go-consul/agent 主对象（watch/replay + register）。
func NewSidecarAgent(bootstrapConfig *conf.BootstrapConfig, services []*grpc.ServiceDesc) (*agent.Agent, error) {
	serviceOpts := &agent.ServiceOptions{
		App:         bootstrapConfig.App,
		Kernel:      bootstrapConfig.Kernel,
		Service:     bootstrapConfig.Service,
		Protocol:    "grpc",
		ServerPort:  bootstrapConfig.ServerPort,
		ManagedPort: bootstrapConfig.ManagedPort,
	}

	// 复制 sidecar 配置后再填充 RawServices，避免直接污染 bootstrap 原对象。
	sidecarCfg := *bootstrapConfig.SidecarAgent
	sidecarCfg.RawServices = services

	return agent.New(serviceOpts, sidecarCfg)
}

// NewSidecarStatusProvider 暴露 sidecar 状态提供者给管理端口复用。
func NewSidecarStatusProvider(a *agent.Agent) SidecarStatusProvider {
	return a
}
