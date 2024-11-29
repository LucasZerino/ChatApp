package models

import (
	"time"
)

// TelegramBot representa a tabela telegram_bots
type TelegramBot struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	AuthKey   string    `gorm:"column:auth_key;not null" json:"auth_key"`
	AccountID uint      `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (TelegramBot) TableName() string {
	return "telegram_bots"
}
