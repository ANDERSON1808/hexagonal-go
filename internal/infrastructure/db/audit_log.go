package db

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID         uint           `gorm:"primaryKey"`
	EntityName string         `gorm:"type:varchar(100);not null"`
	Action     string         `gorm:"type:varchar(50);not null"`
	Timestamp  time.Time      `gorm:"autoCreateTime"`
	Details    string         `gorm:"type:text"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
