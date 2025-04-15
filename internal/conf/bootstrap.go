package conf

import (
	"encoding/json"
	logger "github.com/lhdhtrc/logger-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	task "github.com/lhdhtrc/task-go/pkg"
	"os"
	"path/filepath"
)

type BootstrapConf struct {
	// 运行端口
	Port uint `json:"port"`

	// 应用id
	AppId string `json:"app_id"`
	// 应用密钥
	AppSecret string `json:"app_secret"`

	// 微服务核心组件配置
	Micro *micro.ServiceConf `json:"micro"`
	// 网关配置
	Gateway *micro.GatewayConf `json:"gateway"`
	// 日志组件配置
	Logger *logger.Config `json:"logger"`
	// 任务组件配置
	Task *task.Config `json:"task"`
}

func NewBootstrapConf() *BootstrapConf {
	var bc BootstrapConf

	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file := filepath.Join(cur, "conf", "bootstrap.json")

	var b []byte
	b, err = os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &bc); err != nil {
		panic(err)
	}

	return &bc
}
