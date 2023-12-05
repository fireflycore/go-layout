package base

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type DBTableEntity struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type DBTableUUIDEntity struct {
	ID        uuid.UUID      `json:"id" gorm:"primarykey;type:uuid;default:uuid_generate_v4();"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
