package store

import (
	"go.mongodb.org/mongo-driver/mongo"
	"microservice-go/utils/logger"
)

type _LoggerStoreEntrance struct {
	Table *mongo.Collection
	Func  logger.Interface
	Temp  []string
}
