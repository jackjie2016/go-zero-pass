package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
}
