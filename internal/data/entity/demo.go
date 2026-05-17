package entity

import "github.com/fireflycore/gormx"

// Demo 表示示例业务表实体。
type Demo struct {
	gormx.TableUUID

	// Title 是示例标题。
	Title string `json:"title"`
	// Description 是示例描述。
	Description string `json:"description"`
	// Content 是示例正文内容。
	Content string `json:"content"`
	// Status 表示示例状态。
	Status uint32 `json:"status"`
	// Sort 表示示例排序值。
	Sort uint32 `json:"sort"`

	// UserId 标识所属用户。
	UserId string `json:"user_id" gorm:"type:uuid;index;"`
	// AppId 标识所属应用。
	AppId string `json:"app_id" gorm:"type:uuid;index;"`
	// TenantId 标识所属租户。
	TenantId string `json:"tenant_id" gorm:"type:uuid;index"`
}

// Table 返回示例实体表名。
func (Demo) Table() string {
	return "demo"
}
