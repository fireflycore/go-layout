package server

import (
	"go-layout/internal/conf"

	"github.com/fireflycore/go-consul/agent"
)

// NewSidecarAgent 基于 bootstrap 组装 go-consul/agent 主对象（watch/replay + register）。
func NewSidecarAgent(bootstrapConfig *conf.BootstrapConfig) (*agent.Agent, error) {
	serviceOpts := &agent.ServiceOptions{
		App:         bootstrapConfig.App,
		Kernel:      bootstrapConfig.Kernel,
		Service:     bootstrapConfig.Service,
		Protocol:    "grpc",
		ServerPort:  bootstrapConfig.ServerPort,
		ManagedPort: bootstrapConfig.ManagedPort,
	}

	// 复制 sidecar 配置，避免运行期回填默认值时污染 bootstrap 原对象。
	sidecarCfg := *bootstrapConfig.SidecarAgent
	// gateway.manifest.json 是服务能力唯一输入；GatewayManifestPath 为空时由 go-consul 使用默认路径。
	return agent.New(serviceOpts, sidecarCfg)
}

// NewSidecarStatusProvider 暴露 sidecar 状态提供者给管理端口复用。
func NewSidecarStatusProvider(a *agent.Agent) SidecarStatusProvider {
	return a
}
