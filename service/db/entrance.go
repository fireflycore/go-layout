package db

import (
	"context"
	"github.com/lhdhtrc/microservice-go/db"
	"github.com/lhdhtrc/microservice-go/utils/object"
	pb "microservice-go/dep/protobuf/gen/acme/config/db/v1"
	"microservice-go/store"
	"time"
)

type EntranceEntity struct {
}

func (EntranceEntity) Get(id string) *db.ConfigEntity {
	endpoint := store.Use.Service["/microservice/lhdht/config/db/Add"]

	if len(endpoint) == 0 {
		store.Use.Logger.Warning("the service has no address available")
		return nil
	}

	coon := store.Use.Grpc.Dial(endpoint, &store.Use.Config.Grpc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cli := pb.NewDBConfigServiceClient(coon)
	res, err := cli.Get(ctx, &pb.GetRequest{Id: id})
	if err != nil {
		store.Use.Logger.Warning(err.Error())
		return nil
	}

	var dbConfig db.ConfigEntity
	object.StructConvert(res.Data, &dbConfig)
	return &dbConfig
}
