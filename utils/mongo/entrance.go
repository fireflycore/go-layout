package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"microservice-go/utils/array"
	"microservice-go/utils/time"
)

type Entrance struct {
}

func (Entrance) SetIds(ids []string, query bson.M) {
	if len(ids) != 0 {
		query["_id"] = bson.M{"$in": array.Map[string, primitive.ObjectID](ids, func(index int, item string) primitive.ObjectID {
			hex, _ := primitive.ObjectIDFromHex(item)
			return hex
		})}
	}
}

func (Entrance) TimeFrame(startTime string, endTime string, query bson.M) {
	if startTime != "" && endTime != "" {
		query["created_at"] = bson.M{
			"$gte": time.Entrance{}.StringToTime(startTime),
			"$lt":  time.Entrance{}.StringToTime(endTime),
		}
	}
}

func (Entrance) Paging(page int64, pageSize int64, option *options.FindOptions) {
	if page != 0 {
		option.SetLimit(pageSize)
		option.SetSkip((page - 1) * pageSize)
	}
}
