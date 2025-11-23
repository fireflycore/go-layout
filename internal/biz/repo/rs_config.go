package repo

import (
	"context"
	pb "go-layout/dep/protobuf/gen/acme/config/v1"
)

type ConfigRepo interface {
	GetConfig(ctx context.Context, appId, group, key string) (*pb.Config, error)
}
