package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelAPI representa a tabela channel_api
type ChannelAPI struct {
	ID                   uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID            uint                   `gorm:"column:account_id;not null;index" json:"account_id"`
	WebhookURL           string                 `gorm:"column:webhook_url" json:"webhook_url"`
	CreatedAt            time.Time              `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt            time.Time              `gorm:"column:updated_at;not null" json:"updated_at"`
	Identifier           string                 `gorm:"column:identifier;uniqueIndex" json:"identifier"`
	HMACToken            string                 `gorm:"column:hmac_token;uniqueIndex" json:"hmac_token"`
	HMACMandatory        bool                   `gorm:"column:hmac_mandatory;default:false" json:"hmac_mandatory"`
	AdditionalAttributes JSONB `gorm:"column:additional_attributes;default:'{}'" json:"additional_attributes"`
}

// TableName define o nome da tabela para o GORM
func (ChannelAPI) TableName() string {
	return "channel_api"
}

// CreateIndexes cria os índices para a tabela ChannelAPI
func (ChannelAPI) CreateIndexes(db *gorm.DB) {
	// Criar índices compostos ou únicos
	db.Migrator().CreateIndex(&ChannelAPI{}, "index_channel_api_on_hmac_token")
	db.Migrator().CreateIndex(&ChannelAPI{}, "index_channel_api_on_identifier")
}
