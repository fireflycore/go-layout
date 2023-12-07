package config

import (
	"microservice-go/model/base"
	"microservice-go/model/config"
)

type _MicroConfigEntrance struct {
	Deploy         string                `json:"deploy" yaml:"deploy" mapstructure:"deploy"`                   // Deploy mode (LAN local area network) (NAN not area network)
	GatewayName    string                `json:"gateway_name" yaml:"gateway_name" mapstructure:"gateway_name"` // GatewayName GetawayName The gateway name is used to bind the gateway when registering microservices, facilitating gateway service discovery
	Namespace      string                `json:"namespace" yaml:"namespace" mapstructure:"namespace"`
	Address        base.NetworkAddress   `json:"address" yaml:"address" yaml:"address" mapstructure:"address"`
	RegisterCenter config.DBConfigEntity `json:"register_center" yaml:"register_center" mapstructure:"register_center"` // RegisterCenter config
}
