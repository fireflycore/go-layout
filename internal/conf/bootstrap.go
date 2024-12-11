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
	Server  *ServerConf
	Gateway *GatewayConf

	Logger *logger.Config
	Micro  *micro.ServiceConfig
	Task   *task.Config
}

type GatewayConf struct {
	Network         string `json:"network"`
	OuterNetAddr    string `json:"outer_net_addr"`
	InternalNetAddr string `json:"internal_net_addr"`
}

type ServerConf struct {
	GrpcPort uint `json:"grpc_port"`
}

type LoggerConf struct {
	Remote  bool `json:"remote"`
	Console bool `json:"console"`
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
