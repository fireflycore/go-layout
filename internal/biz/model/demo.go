package model

import (
	gorm "github.com/lhdhtrc/gorm/pkg"
	"gorm.io/datatypes"
)

type Demo struct {
	gorm.TableUUID

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`

	UserId   datatypes.UUID
	AppId    datatypes.UUID
	TenantId datatypes.UUID
}

type CreateDemo struct {
	Title       string
	Description string
	Content     string
	Status      uint32
	Sort        uint32

	UserId   datatypes.UUID
	AppId    datatypes.UUID
	TenantId datatypes.UUID
}

type UpdateDemo struct {
	Title       string
	Description string
	Content     string
	Status      uint32
	Sort        uint32
}
