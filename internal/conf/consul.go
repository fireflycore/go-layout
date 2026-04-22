package conf

import "github.com/fireflycore/go-consul"

func NewConsulConf(utils *Utils) (*consul.Config, error) {
	var dst consul.Config
	if err := utils.LoadJSONConfig(utils.GetConfigFilePath("consul.json"), &dst); err != nil {
		return nil, err
	}
	return &dst, nil
}
