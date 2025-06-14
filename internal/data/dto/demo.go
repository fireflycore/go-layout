package dto

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"
)

// goverter:converter
// goverter:output:package :data
// goverter:output:file @cwd/dep/goverter/data/demo.dto.go
// goverter:extend StringToUUID UUIDToString UUIDPtrToString TimeToString
type DemoConverter interface {
	// goverter:ignore TableUUID UserId AppId
	ToCreate(row *pb.CreateRequest) *entity.Demo
	// goverter:ignore TableUUID UserId AppId
	ToUpdate(row *pb.UpdateRequest) *entity.Demo

	// goverter:ignore state unknownFields sizeCache Author
	// goverter:map TableUUID.ID Id
	// goverter:map TableUUID.CreatedAt CreatedAt
	// goverter:map TableUUID.UpdatedAt UpdatedAt
	ToRow(row *entity.Demo) *pb.Demo
	ToRaw(raw []*entity.Demo) []*pb.Demo
}
