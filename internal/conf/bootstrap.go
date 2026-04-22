package conf

import (
	"net"
	"strconv"

	microConfig "github.com/fireflycore/go-micro/config"
	"github.com/fireflycore/go-micro/constant"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-micro/telemetry"
	"github.com/fireflycore/go-utils/network"
	"github.com/google/uuid"
)

// SidecarAgentConf 描述本地 sidecar-agent 的接入参数。
type SidecarAgentConf struct {
	BaseURL     string `json:"base_url"`
	GracePeriod string `json:"grace_period"`
}

// BootstrapConf 是模板库统一复用的服务启动配置。
type BootstrapConf struct {
	Env string `json:"env"`

	Port           uint `json:"port"`
	ServerPort     uint `json:"server_port"`
	ManagementPort uint `json:"management_port"`

	AppId     string `json:"app_id"`
	AppName   string `json:"app_name"`
	AppSecret string `json:"app_secret"`
	Version   string `json:"version"`

	ServiceName      string `json:"service_name"`
	ServiceNamespace string `json:"service_namespace"`
	ServiceDNS       string `json:"service_dns"`

	SidecarAgent *SidecarAgentConf `json:"sidecar_agent"`
	Logger       *logger.Conf      `json:"logger"`
	Telemetry    *telemetry.Conf   `json:"telemetry"`

	ServiceInstanceId string        `json:"-"`
	SystemHostInfo    *sys.HostInfo `json:"-"`
}

// NewBootstrapConf 加载并补齐服务启动配置，确保模板默认即可接入 sidecar 托管模型。
func NewBootstrapConf(utils *Utils, hostInfo *sys.HostInfo) *BootstrapConf {
	var bc BootstrapConf

	filePath := utils.GetConfigFilePath("bootstrap.json")
	if err := utils.LoadJSONConfig(filePath, &bc); err != nil {
		panic(err)
	}

	if bc.ServerPort == 0 {
		bc.ServerPort = bc.Port
	}
	if bc.ManagementPort == 0 && bc.ServerPort != 0 {
		bc.ManagementPort = bc.ServerPort + 1
	}
	if bc.ServiceName == "" {
		bc.ServiceName = "go-layout"
	}
	if bc.ServiceNamespace == "" {
		bc.ServiceNamespace = "default"
	}
	if bc.ServiceDNS == "" && bc.ServiceName != "" {
		bc.ServiceDNS = bc.ServiceName + "." + bc.ServiceNamespace + ".svc.cluster.local"
	}
	if bc.SidecarAgent == nil {
		bc.SidecarAgent = &SidecarAgentConf{GracePeriod: "20s"}
	}
	if bc.Logger == nil {
		bc.Logger = &logger.Conf{}
	}
	if bc.Telemetry == nil {
		bc.Telemetry = &telemetry.Conf{}
	}

	bc.ServiceInstanceId = uuid.Must(uuid.NewV7()).String()
	bc.SystemHostInfo = hostInfo

	return &bc
}

// NewBootstrapConfImpl 把具体配置适配为 go-micro 统一的 BootstrapConfig 接口。
func NewBootstrapConfImpl(bootstrapConf *BootstrapConf) microConfig.BootstrapConfig {
	return bootstrapConf
}

func (bc *BootstrapConf) GetAppId() string { return bc.AppId }

func (bc *BootstrapConf) GetAppSecret() string { return bc.AppSecret }

func (bc *BootstrapConf) GetAppName() string { return bc.AppName }

func (bc *BootstrapConf) GetAppVersion() string { return bc.Version }

func (bc *BootstrapConf) GetServiceEndpoint() string {
	return net.JoinHostPort(network.GetInternalNetworkIp(), strconv.FormatUint(uint64(bc.ServerPort), 10))
}

func (bc *BootstrapConf) GetServiceAuthToken() string { return constant.InvokeServiceAuthToken }

func (bc *BootstrapConf) GetServiceNamespace() string { return bc.ServiceNamespace }

func (bc *BootstrapConf) GetServiceInstanceId() string { return bc.ServiceInstanceId }

func (bc *BootstrapConf) GetSystemName() string { return bc.SystemHostInfo.OS }

func (bc *BootstrapConf) GetSystemType() uint32 { return bc.SystemHostInfo.GetSystemType() }

func (bc *BootstrapConf) GetSystemVersion() string { return bc.SystemHostInfo.PlatformVersion }

func (bc *BootstrapConf) GetServerPort() uint { return bc.ServerPort }

func (bc *BootstrapConf) GetManagementPort() uint { return bc.ManagementPort }

func (bc *BootstrapConf) GetLoggerConsole() bool { return bc.Logger.GetLoggerConsole() }

func (bc *BootstrapConf) GetLoggerRemote() bool { return bc.Logger.GetLoggerRemote() }

func (bc *BootstrapConf) GetOtelEndpoint() string { return bc.Telemetry.GetOtelEndpoint() }

func (bc *BootstrapConf) GetOtelInsecure() bool { return bc.Telemetry.GetOtelInsecure() }

func (bc *BootstrapConf) GetOtelTraces() bool { return bc.Telemetry.GetOtelTraces() }

func (bc *BootstrapConf) GetOtelMetrics() bool { return bc.Telemetry.GetOtelMetrics() }

func (bc *BootstrapConf) GetOtelLogs() bool { return bc.Telemetry.GetOtelLogs() }

func (bc *BootstrapConf) GetServiceName() string { return bc.ServiceName }

func (bc *BootstrapConf) GetServiceDNS() string { return bc.ServiceDNS }

func (bc *BootstrapConf) GetSidecarAgentBaseURL() string {
	if bc.SidecarAgent == nil {
		return ""
	}
	return bc.SidecarAgent.BaseURL
}

func (bc *BootstrapConf) GetSidecarGracePeriod() string {
	if bc.SidecarAgent == nil || bc.SidecarAgent.GracePeriod == "" {
		return "20s"
	}
	return bc.SidecarAgent.GracePeriod
}
