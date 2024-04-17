package config

type SystemConfigEntity struct {
	AppId         string `json:"app_id" bson:"app_id" yaml:"app_id" mapstructure:"app_id"`
	RunPort       string `json:"run_port" bson:"run_port" yaml:"run_port" mapstructure:"run_port"`
	Deploy        bool   `json:"deploy" bson:"deploy" yaml:"deploy" mapstructure:"deploy"`
	DeployAddress bool   `json:"deploy_address" bson:"deploy_address" yaml:"deploy_address" mapstructure:"deploy_address"` // outside:inside
}
