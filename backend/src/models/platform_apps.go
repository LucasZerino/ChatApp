package models

import (
	"time"
)

// PlatformApp representa a tabela platform_apps
type PlatformApp struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (PlatformApp) TableName() string {
	return "platform_apps"
}
