package conf

import (
	"net"
	"strconv"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/app"
	"github.com/fireflycore/go-micro/kernel"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/service"
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-micro/telemetry"
)

// BootstrapConfig 是服务启动阶段加载的一次性引导配置。
type BootstrapConfig struct {
	App       app.Config       `json:"app"`
	Kernel    kernel.Config    `json:"kernel"`
	Logger    logger.Config    `json:"logger"`
	Service   service.Config   `json:"service"`
	Telemetry telemetry.Config `json:"telemetry"`

	ServerPort     uint   `json:"server_port"`
	ManagedPort    uint   `json:"managed_port"`

	SidecarAgent   *agent.SidecarAgentConfig `json:"sidecar_agent"`
	SystemHostInfo *sys.HostInfo             `json:"-"`
}

// NewBootstrapConfig 加载并补齐服务启动配置。
func NewBootstrapConfig(utils *Utils, hostInfo *sys.HostInfo) *BootstrapConfig {
	var bc BootstrapConfig

	filePath := utils.GetConfigFilePath("bootstrap.json")
	if err := utils.LoadJSONConfig(filePath, &bc); err != nil {
		panic(err)
	}

	if err := bc.App.Bootstrap(); err != nil {
		panic(err)
	}
	bc.Kernel.Bootstrap()
	if err := bc.Service.Bootstrap(); err != nil {
		panic(err)
	}
	if bc.ServerPort == 0 {
		bc.ServerPort = bc.Service.Port
	}
	if bc.ManagedPort == 0 && bc.ServerPort != 0 {
		bc.ManagedPort = bc.ServerPort + 1
	}
	if bc.SidecarAgent == nil {
		sidecarAgent := agent.DefaultSidecarAgentConfig("")
		bc.SidecarAgent = &sidecarAgent
	} else {
		sidecarAgent := agent.DefaultSidecarAgentConfig(bc.SidecarAgent.BaseURL)
		if bc.SidecarAgent.WatchURL != "" {
			sidecarAgent.WatchURL = bc.SidecarAgent.WatchURL
		}
		if bc.SidecarAgent.GracePeriod != "" {
			sidecarAgent.GracePeriod = bc.SidecarAgent.GracePeriod
		}
		if bc.SidecarAgent.RequestTimeout > 0 {
			sidecarAgent.RequestTimeout = bc.SidecarAgent.RequestTimeout
		}
		if bc.SidecarAgent.ReconnectInterval > 0 {
			sidecarAgent.ReconnectInterval = bc.SidecarAgent.ReconnectInterval
		}
		bc.SidecarAgent = &sidecarAgent
	}
	if bc.SidecarAgent.GracePeriod == "" {
		bc.SidecarAgent.GracePeriod = "20s"
	}

	bc.SystemHostInfo = hostInfo
	return &bc
}

func NewLoggerConfig(bootstrapConfig *BootstrapConfig) *logger.Config {
	return &bootstrapConfig.Logger
}

func (bc *BootstrapConfig) ServiceEndpoint() string {
	return net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bc.ServerPort), 10))
}

func (bc *BootstrapConfig) ManagementEndpoint() string {
	return net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bc.ManagedPort), 10))
}

func (bc *BootstrapConfig) SidecarAgentConfig() agent.SidecarAgentConfig {
	if bc == nil || bc.SidecarAgent == nil {
		return agent.DefaultSidecarAgentConfig("")
	}
	return *bc.SidecarAgent
}
