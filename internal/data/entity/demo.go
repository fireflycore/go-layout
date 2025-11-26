package entity

import (
	gorm "github.com/lhdhtrc/gorm/pkg"
)

type Demo struct {
	gorm.TableUUID

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`

	UserId   string `json:"user_id" gorm:"type:uuid;index;"`
	AppId    string `json:"app_id" gorm:"type:uuid;index;"`
	TenantId string `json:"tenant_id" gorm:"type:uuid;index"`
}

func (Demo) Table() string {
	return "demo"
}
