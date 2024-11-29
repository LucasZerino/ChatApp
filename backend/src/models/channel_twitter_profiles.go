package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelTwitterProfile representa a tabela channel_twitter_profiles
type ChannelTwitterProfile struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProfileID                 string    `gorm:"column:profile_id;not null;uniqueIndex:idx_account_id_and_profile_id" json:"profile_id"`
	TwitterAccessToken        string    `gorm:"column:twitter_access_token;not null" json:"twitter_access_token"`
	TwitterAccessTokenSecret  string    `gorm:"column:twitter_access_token_secret;not null" json:"twitter_access_token_secret"`
	AccountID                 uint      `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt                 time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	TweetsEnabled             bool      `gorm:"column:tweets_enabled;default:true" json:"tweets_enabled"`
}

// TableName define o nome da tabela para o GORM
func (ChannelTwitterProfile) TableName() string {
	return "channel_twitter_profiles"
}

// CreateIndexes cria os índices para a tabela ChannelTwitterProfile
func (ChannelTwitterProfile) CreateIndexes(db *gorm.DB) {
	// Criar índice único para os campos account_id e profile_id
	db.Migrator().CreateIndex(&ChannelTwitterProfile{}, "index_channel_twitter_profiles_on_account_id_and_profile_id")
}
