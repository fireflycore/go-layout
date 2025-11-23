package convert

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/model"
	"go-layout/internal/data/entity"
)

// goverter:converter
// goverter:output:package :dto
// goverter:output:file @cwd/dep/dto/demo.go
// goverter:extend StringToUUID UUIDToString UUIDPtrToString TimeToString
type DemoConverter interface {
	// goverter:ignore TableUUID UserId AppId TenantId
	ToCreate(row *pb.CreateDemoRequest) *entity.Demo
	ToUpdate(row *pb.UpdateDemoRequest) *model.UpdateDemo

	// goverter:ignore state unknownFields sizeCache Author
	// goverter:map TableUUID.ID Id
	// goverter:map TableUUID.CreatedAt CreatedAt
	// goverter:map TableUUID.UpdatedAt UpdatedAt
	ToRow(row *entity.Demo) *pb.Demo
	ToRaw(raw []*entity.Demo) []*pb.Demo
}
