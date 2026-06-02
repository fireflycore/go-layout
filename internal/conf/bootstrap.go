package conf

import (
	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/app"
	"github.com/fireflycore/go-micro/authz"
	"github.com/fireflycore/go-micro/kernel"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/service"
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-micro/telemetry"
)

// BootstrapConfig 描述服务启动期静态配置。
type BootstrapConfig struct {
	// App 应用配置
	App app.Config `json:"app"`
	// Kernel 内核配置
	Kernel kernel.Config `json:"kernel"`
	// Logger 日志配置
	Logger logger.Config `json:"logger"`
	// Service 服务配置
	Service service.Config `json:"service"`
	// Telemetry 可观测性配置
	Telemetry telemetry.Config `json:"telemetry"`

	// ServerPort 服务端口
	ServerPort uint `json:"server_port"`
	// ManagePort 管理端口
	ManagedPort uint `json:"managed_port"`

	// SidecarAgent 保存 sidecar-agent 接管服务生命周期所需配置。
	SidecarAgent *agent.SidecarAgentConfig `json:"sidecar_agent"`

	// AuthzVerification 保存服务侧本地验签 x-firefly-authz-sign 所需配置。
	//
	// 为空表示当前模板只解析普通 metadata，不启用本地验签；
	// 非空表示服务启动时必须能加载 authz Ed25519 公钥。
	AuthzVerification *authz.VerificationConfig `json:"authz_verification"`

	// SystemHostInfo 保存宿主机信息，启动后由代码注入，不从配置文件反序列化。
	SystemHostInfo *sys.HostInfo `json:"-"`
}

// NewBootstrapConfig 读取并初始化服务启动配置。
func NewBootstrapConfig(utils *Utils, hostInfo *sys.HostInfo) *BootstrapConfig {
	var bc BootstrapConfig
	// bootstrap.json 只承载启动期静态配置，不走运行时配置中心。
	filePath := utils.GetConfigFilePath("bootstrap.json")
	if err := utils.LoadJSONConfig(filePath, &bc); err != nil {
		panic(err)
	}

	// App / Service 的 Bootstrap 会补齐实例标识、服务默认值等运行必要字段。
	if err := bc.App.Bootstrap(); err != nil {
		panic(err)
	}
	bc.Kernel.Bootstrap()
	if err := bc.Service.Bootstrap(); err != nil {
		panic(err)
	}
	// 宿主机信息在启动期注入，供日志与注册链路复用。
	bc.SystemHostInfo = hostInfo

	return &bc
}

// NewLoggerConfig 暴露日志配置给 logger provider 复用。
func NewLoggerConfig(bootstrapConfig *BootstrapConfig) *logger.Config {
	return &bootstrapConfig.Logger
}
