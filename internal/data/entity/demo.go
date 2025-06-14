package entity

import gorm "github.com/lhdhtrc/gorm/pkg"

type Demo struct {
	gorm.TableUUID

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`

	AppId  gorm.BinUUID `json:"app_id" gorm:"index;"`
	UserId gorm.BinUUID `json:"user_id" gorm:"index;"`
}

func (Demo) Table() string {
	return "demo"
}
