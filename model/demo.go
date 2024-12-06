package model

import gorm "github.com/lhdhtrc/gorm/pkg"

type DemoEntity struct {
	gorm.TableUUID

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`

	AppId     string `json:"app_id" gorm:"size:36;index;"`
	AccountId string `json:"account_id" gorm:"size:36;index;"`
}
