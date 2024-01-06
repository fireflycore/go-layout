package logger

import (
	"context"
	"github.com/lhdhtrc/microservice-go/logger"
	pb "microservice-go/dep/protobuf/gen/acme/logger/v1"
	"microservice-go/store"
	"time"
)

type EntranceEntity struct {
}

func (EntranceEntity) Add(entity logger.Entity) {
	endpoint := store.Use.Service["/microservice/lhdht/logger/Add"]

	if len(endpoint) == 0 {
		return
	}

	coon := store.Use.Grpc.Dial(endpoint, &store.Use.Config.Grpc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cli := pb.NewLoggerServiceClient(coon)

	if len(store.Use.Logger.Temp) != 0 {
		for _, log := range store.Use.Logger.Temp {
			if _, err := cli.Add(ctx, &pb.AddRequest{AppId: store.Use.Config.System.AppId, Level: log.Level, Message: log.Message}); err != nil {
				store.Use.Logger.Warning(err.Error())
				return
			}
		}
		store.Use.Logger.Temp = nil
	}

	_, err := cli.Add(ctx, &pb.AddRequest{
		AppId:   store.Use.Config.System.AppId,
		Level:   entity.Level,
		Message: entity.Message,
	})
	if err != nil {
		store.Use.Logger.Warning(err.Error())
		return
	}
}
