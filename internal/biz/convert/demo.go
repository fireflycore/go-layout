package convert

import (
	pb "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/data/entity"
)

// goverter:converter
// goverter:output:package :dto
// goverter:output:file @cwd/dep/dto/demo.go
// DemoConvert 定义 Demo 在 Proto 与实体之间的转换规则。
type DemoConvert interface {
	// goverter:ignore TableUUID UserId AppId TenantId
	// ToCreate 把创建请求转换成 Demo 实体。
	ToCreate(row *pb.CreateDemoRequest) *entity.Demo
}
