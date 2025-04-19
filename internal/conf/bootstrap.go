package conf

import (
	"fmt"
	logger "github.com/lhdhtrc/logger-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
)

type BootstrapConf struct {
	// 环境
	Env string `json:"env"`
	// 运行端口
	Port uint `json:"port"`

	// 应用id
	AppId string `json:"app_id"`
	// 应用密钥
	AppSecret string `json:"app_secret"`

	// 版本号
	Version string `json:"version"`

	// 加载配置模式: local 为本地加载 remote 为远程加载
	LoadConfMode string `json:"load_conf_mode"`

	// 微服务核心组件配置
	Micro *micro.ServiceConf `json:"micro"`
	// 网关配置
	Gateway *micro.GatewayConf `json:"gateway"`
	// 日志组件配置
	Logger *logger.Config `json:"logger"`
}

func NewBootstrapConf() *BootstrapConf {
	var bc BootstrapConf

	filePath := getConfigFilePath("bootstrap.json")

	if err := loadJSONConfig(filePath, &bc); err != nil {
		panic(err)
	}

	bc.Micro.Network.Internal = fmt.Sprintf("%s:%d", micro.GetInternalNetworkIp(), bc.Port)

	return &bc
}
