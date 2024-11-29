package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelTelegram representa a tabela channel_telegram
type ChannelTelegram struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	BotName     string    `gorm:"column:bot_name" json:"bot_name"`
	AccountID   uint      `gorm:"column:account_id;not null" json:"account_id"`
	BotToken    string    `gorm:"column:bot_token;not null;uniqueIndex:idx_bot_token" json:"bot_token"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (ChannelTelegram) TableName() string {
	return "channel_telegram"
}

// CreateIndexes cria os índices para a tabela ChannelTelegram
func (ChannelTelegram) CreateIndexes(db *gorm.DB) {
	// Criar índice único para o campo bot_token
	db.Migrator().CreateIndex(&ChannelTelegram{}, "index_channel_telegram_on_bot_token")
}
