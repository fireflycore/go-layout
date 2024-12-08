package conf

import (
	"encoding/json"
	micro "github.com/lhdhtrc/micro-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
	"os"
	"path/filepath"
)

type BootstrapConf struct {
	Server  *ServerConf
	Logger  *LoggerConf
	Gateway *GatewayConf

	Micro *micro.ServiceConfig
	Task  *task.Config
}

type GatewayConf struct {
	Network    string `json:"network"`
	OuterAddr  string `json:"outer_addr"`
	InsideAddr string `json:"inside_addr"`
}

type ServerConf struct {
	GrpcPort string `json:"grpc_port"`
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
