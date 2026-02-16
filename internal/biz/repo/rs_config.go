package repo

import (
	pb "go-layout/dep/protobuf/gen/acme/config/v1"
)

type ConfigRepo interface {
	GetConfig(appId, group, key string) (*pb.Config, error)
}
