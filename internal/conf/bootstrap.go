package conf

import (
	"net"
	"strconv"

	"github.com/fireflycore/go-micro/conf"
	"github.com/fireflycore/go-micro/constant"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/registry"
	"github.com/fireflycore/go-micro/sys"
	"github.com/fireflycore/go-utils/network"
)

type BootstrapConf struct {
	// 环境
	Env string `json:"env"`
	// 运行端口
	Port uint `json:"port"`

	// 应用id
	AppId string `json:"app_id"`
	// 应用名称
	AppName string `json:"app_name"`
	// 应用密钥
	AppSecret string `json:"app_secret"`

	// 版本号
	Version string `json:"version"`

	// 加载配置模式: local 为本地加载 remote 为远程加载
	LoadConfMode string `json:"load_conf_mode"`

	// 微服务核心组件配置
	Micro *registry.ServiceConf `json:"micro"`
	// 网关配置
	Gateway *registry.GatewayConf `json:"gateway"`
	// 日志组件配置
	Logger *logger.Conf `json:"logger"`

	// 系统主机信息
	SystemHostInfo *sys.HostInfo `json:"-"`
}

func NewBootstrapConf(utils *Utils, hostInfo *sys.HostInfo) *BootstrapConf {
	var bc BootstrapConf

	filePath := utils.GetConfigFilePath("bootstrap.json")

	if err := utils.LoadJSONConfig(filePath, &bc); err != nil {
		panic(err)
	}

	bc.Micro.Network.Internal = net.JoinHostPort(network.GetInternalNetworkIp(), strconv.FormatUint(uint64(bc.Port), 10))
	bc.SystemHostInfo = hostInfo

	return &bc
}

func NewBootstrapConfImpl(bootstrapConf *BootstrapConf) conf.BootstrapConf {
	return bootstrapConf
}

func (bc *BootstrapConf) GetAppId() string {
	return bc.AppId
}

func (bc *BootstrapConf) GetAppName() string {
	return bc.AppName
}

func (bc *BootstrapConf) GetAppVersion() string {
	return bc.Version
}

func (bc *BootstrapConf) GetServiceEndpoint() string {
	return bc.Micro.Network.Internal
}

func (bc *BootstrapConf) GetServiceAuthToken() string {
	return constant.InvokeServiceAuthToken
}

func (bc *BootstrapConf) GetSystemType() uint32 {
	return bc.SystemHostInfo.GetSystemType()
}

func (bc *BootstrapConf) GetSystemName() string {
	return bc.SystemHostInfo.OS
}

func (bc *BootstrapConf) GetSystemVersion() string {
	return bc.SystemHostInfo.PlatformVersion
}

func (bc *BootstrapConf) GetGatewayEndpoint() string {
	var addr string
	if bc.Gateway.Network.SN == bc.Micro.Network.SN {
		addr = bc.Gateway.Network.Internal
	} else {
		addr = bc.Gateway.Network.External
	}
	return addr
}
