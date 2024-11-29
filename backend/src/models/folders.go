package models

import (
	"time"
)

// Folder representa a tabela folders
type Folder struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID  uint      `gorm:"column:account_id;not null" json:"account_id"`
	CategoryID uint     `gorm:"column:category_id;not null" json:"category_id"`
	Name       string    `gorm:"column:name" json:"name"`
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Folder) TableName() string {
	return "folders"
}
