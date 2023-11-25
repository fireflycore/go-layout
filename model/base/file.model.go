package base

type FileInfoEntity struct {
	Name string `json:"name" bson:"name" yaml:"name" mapstructure:"name"`
	Type string `json:"type" bson:"type" yaml:"type" mapstructure:"type"`
	Url  string `json:"url" bson:"url" yaml:"url" mapstructure:"url"`
	Size string `json:"size" bson:"size" yaml:"size" mapstructure:"size"`
}
