package base

type DescriptionEntity struct {
	Description string `json:"description" bson:"description" yaml:"description" mapstructure:"description"`
	Remark      string `json:"remark" bson:"remark" yaml:"remark" mapstructure:"remark"`
	State       uint   `json:"state" bson:"state" yaml:"state" mapstructure:"state"`
}
