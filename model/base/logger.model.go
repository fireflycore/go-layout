package base

type LoggerEntity struct {
	Level   uint   `json:"level" bson:"level"`
	Message string `json:"message" bson:"message"`
	AppId   string `json:"app_id" bson:"app_id"`
}
