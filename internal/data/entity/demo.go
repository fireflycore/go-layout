package entity

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

	UserId   datatypes.UUID `json:"user_id" gorm:"index;"`
	AppId    datatypes.UUID `json:"app_id" gorm:"index;"`
	TenantId datatypes.UUID `json:"tenant_id" gorm:"index"`
}

func (Demo) Table() string {
	return "demo"
}
