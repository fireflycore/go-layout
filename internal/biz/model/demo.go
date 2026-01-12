package model

import (
	gorm "github.com/fireflycore/gormx"
)

type Demo struct {
	gorm.TableUUID

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`

	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`

	AppId    string `json:"app_id"`
	TenantId string `json:"tenant_id"`
}
