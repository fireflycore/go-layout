package conf

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	configCenter "go-layout/dep/protobuf/gen/acme/config/v1"
)

// Loader 定义配置加载的统一接口
type Loader interface {
	// LoadLocal 从本地文件系统加载配置
	LoadLocal(dc *DataConf) error
	// LoadRemote 从远程配置中心加载配置
	LoadRemote(dc *DataConf, bc *BootstrapConf, cc configCenter.ConfigCenterServiceClient) error
}

type DataConf struct {
	Etcd *etcd.Config // Etcd连接配置
}

func NewDataConf(bc *BootstrapConf, cc configCenter.ConfigCenterServiceClient) *DataConf {
	dc := new(DataConf)

	// 注册配置加载器
	loaders := []Loader{
		&EtcdLoader{},
	}

	// 执行所有配置加载
	for _, loader := range loaders {
		switch bc.LoadConfMode {
		case "local":
			// 从本地加载
			if err := loader.LoadLocal(dc); err != nil {
				panic(err)
			}
		case "remote":
			// 从配置中心加载
			if err := loader.LoadRemote(dc, bc, cc); err != nil {
				panic(err)
			}
		default:
			panic("not found load conf mode")
		}
	}

	// 环境特定覆盖
	if bc.Env == "dev" {
		dc.Etcd.Endpoint = []string{"119.45.227.16:10106"}
	}

	return dc
}
