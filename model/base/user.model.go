package base

type AccountPasswordEntity struct {
	Account  string `json:"account" bson:"account"  yaml:"account" mapstructure:"account"`
	Password string `json:"password" bson:"password" yaml:"password" mapstructure:"password"`
}

type SexAgeEntity struct {
	Sex uint `json:"sex" bson:"sex" yaml:"sex" mapstructure:"sex"`
	Age uint `json:"age" bson:"age" yaml:"age" mapstructure:"age"`
}

type IdCardInfoEntity struct {
	IdCardNumber string `json:"id_card_number" bson:"id_card_number" yaml:"id_card_number" mapstructure:"id_card_number"`
	IdCardType   uint   `json:"id_card_type" bson:"id_card_type" yaml:"id_card_type" mapstructure:"id_card_type"`
	IdCardState  uint   `json:"id_card_state" bson:"id_card_state" yaml:"id_card_state" mapstructure:"id_card_state"`
}
