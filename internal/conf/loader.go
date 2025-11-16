package conf

import config "go-layout/depend/protobuf/gen/acme/config/v1"

// Loader 定义配置加载的统一接口
type Loader interface {
	// Local 从本地文件系统加载配置
	Local(dc *Conf) error
	// Remote 从远程配置中心加载配置
	Remote(dc *Conf, bc *BootstrapConf, cc config.ConfigServiceClient) error
}
