package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelWhatsapp representa a tabela channel_whatsapp
type ChannelWhatsapp struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID                 uint      `gorm:"column:account_id;not null" json:"account_id"`
	PhoneNumber              string    `gorm:"column:phone_number;not null" json:"phone_number"`
	Provider                 string    `gorm:"column:provider;default:'default'" json:"provider"`
	ProviderConfig           string    `gorm:"column:provider_config;type:jsonb;default:'{}'" json:"provider_config"`
	CreatedAt                time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	MessageTemplates         string    `gorm:"column:message_templates;type:jsonb;default:'{}'" json:"message_templates"`
	MessageTemplatesLastUpdated time.Time `gorm:"column:message_templates_last_updated" json:"message_templates_last_updated"`
}

// TableName define o nome da tabela para o GORM
func (ChannelWhatsapp) TableName() string {
	return "channel_whatsapp"
}

// CreateIndexes cria os índices para a tabela ChannelWhatsapp
func (ChannelWhatsapp) CreateIndexes(db *gorm.DB) {
	// Criar índice único para o campo phone_number
	db.Migrator().CreateIndex(&ChannelWhatsapp{}, "index_channel_whatsapp_on_phone_number")
}
