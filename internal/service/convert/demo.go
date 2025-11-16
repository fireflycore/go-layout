package convert

import (
	pb "go-layout/depend/protobuf/gen/acme/demo/v1"
	"go-layout/internal/biz/model"
)

// goverter:converter
// goverter:output:package :service
// goverter:output:file @cwd/depend/convert/service/demo.go
// goverter:extend StringToUUID UUIDToString UUIDPtrToString TimeToString
type DemoConverter interface {
	// goverter:ignore UserId AppId TenantId
	ToCreate(row *pb.CreateDemoRequest) *model.CreateDemo
	ToUpdate(row *pb.UpdateDemoRequest) *model.UpdateDemo

	// goverter:ignore state unknownFields sizeCache Author
	// goverter:map TableUUID.ID Id
	// goverter:map TableUUID.CreatedAt CreatedAt
	// goverter:map TableUUID.UpdatedAt UpdatedAt
	ToRow(row *model.Demo) *pb.Demo
	ToRaw(raw []*model.Demo) []*pb.Demo
}
