package conf

import "github.com/fireflycore/go-consul"

// NewConsulConfig 读取 Consul 客户端配置。
func NewConsulConfig(utils *Utils) (*consul.Config, error) {
	var dst consul.Config
	// 这里只负责本地静态文件加载，不在 conf 层做额外默认值修补。
	if err := utils.LoadJSONConfig(utils.GetConfigFilePath("consul.json"), &dst); err != nil {
		return nil, err
	}
	return &dst, nil
}
