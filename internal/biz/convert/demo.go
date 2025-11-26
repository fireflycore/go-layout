package convert

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"
)

// goverter:converter
// goverter:output:package :dto
// goverter:output:file @cwd/dep/dto/demo.go
type DemoConvert interface {
	// goverter:ignore TableUUID UserId AppId TenantId
	ToCreate(row *pb.CreateDemoRequest) *entity.Demo
}
