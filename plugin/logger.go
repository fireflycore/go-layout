package plugin

import (
	"context"
	"encoding/json"
	pb "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/store"
	"time"
)

func InstallServerLogger(b []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var row pb.AddRequest
	_ = json.Unmarshal(b, &row)
	row.AppId = store.Use.Config.System.AppId

	_, _ = store.Use.RemoteService.ServerLogger.Add(ctx, &row)
}
