package config

type SystemConfigEntity struct {
	AppId string `json:"app_id" bson:"app_id" yaml:"app_id" mapstructure:"app_id"`
}
