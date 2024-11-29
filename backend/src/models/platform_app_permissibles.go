package models

import (
	"time"
)

// PlatformAppPermissible representa a tabela platform_app_permissibles
type PlatformAppPermissible struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PlatformAppID      uint      `gorm:"column:platform_app_id;not null" json:"platform_app_id"`
	PermissibleType    string    `gorm:"column:permissible_type;not null" json:"permissible_type"`
	PermissibleID      uint      `gorm:"column:permissible_id;not null" json:"permissible_id"`
	CreatedAt          time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (PlatformAppPermissible) TableName() string {
	return "platform_app_permissibles"
}
