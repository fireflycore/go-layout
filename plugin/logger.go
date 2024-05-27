package plugin

import (
	"context"
	"encoding/json"
	pb "microservice-go/dep/protobuf/gen/acme/logger/server/v1"
	"microservice-go/store"
	"time"
)

func SetupRemoteLogger(b []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// base example
	var row pb.AddRequest
	_ = json.Unmarshal(b, &row)
	row.AppId = store.Use.Config.System.AppId

	_, _ = store.Use.RemoteService.LoggerServer.Add(ctx, &row)
}
