package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelTwilioSMS representa a tabela channel_twilio_sms
type ChannelTwilioSMS struct {
	ID                     uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PhoneNumber            string    `gorm:"column:phone_number;not null;uniqueIndex:idx_phone_number" json:"phone_number"`
	AuthToken             string    `gorm:"column:auth_token;not null" json:"auth_token"`
	AccountSID            string    `gorm:"column:account_sid;not null;uniqueIndex:idx_account_sid_and_phone_number" json:"account_sid"`
	AccountID             uint      `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt             time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Medium                int       `gorm:"column:medium;default:0" json:"medium"`
	MessagingServiceSID   string    `gorm:"column:messaging_service_sid" json:"messaging_service_sid"`
	APIKeySID             string    `gorm:"column:api_key_sid" json:"api_key_sid"`
}

// TableName define o nome da tabela para o GORM
func (ChannelTwilioSMS) TableName() string {
	return "channel_twilio_sms"
}

// CreateIndexes cria os índices para a tabela ChannelTwilioSMS
func (ChannelTwilioSMS) CreateIndexes(db *gorm.DB) {
	// Criar índice único para os campos account_sid e phone_number
	db.Migrator().CreateIndex(&ChannelTwilioSMS{}, "index_channel_twilio_sms_on_account_sid_and_phone_number")
	// Criar índice único para o campo messaging_service_sid
	db.Migrator().CreateIndex(&ChannelTwilioSMS{}, "index_channel_twilio_sms_on_messaging_service_sid")
	// Criar índice único para o campo phone_number
	db.Migrator().CreateIndex(&ChannelTwilioSMS{}, "index_channel_twilio_sms_on_phone_number")
}
