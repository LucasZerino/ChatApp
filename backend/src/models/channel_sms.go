package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelSMS representa a tabela channel_sms
type ChannelSMS struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID       uint      `gorm:"column:account_id;not null" json:"account_id"`
	PhoneNumber     string    `gorm:"column:phone_number;not null;uniqueIndex:idx_phone_number" json:"phone_number"`
	Provider        string    `gorm:"column:provider;default:'default'" json:"provider"`
	ProviderConfig  string    `gorm:"column:provider_config;type:jsonb;default:'{}'" json:"provider_config"`
	CreatedAt       time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (ChannelSMS) TableName() string {
	return "channel_sms"
}

// CreateIndexes cria os índices para a tabela ChannelSMS
func (ChannelSMS) CreateIndexes(db *gorm.DB) {
	// Criar índice único para o campo phone_number
	db.Migrator().CreateIndex(&ChannelSMS{}, "index_channel_sms_on_phone_number")
}
