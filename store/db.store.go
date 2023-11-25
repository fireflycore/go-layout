package store

import "go.mongodb.org/mongo-driver/mongo"

type _DBStoreEntrance struct {
	Mongo map[string]*mongo.Database
}
