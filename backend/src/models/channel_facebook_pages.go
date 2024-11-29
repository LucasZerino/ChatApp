package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelFacebookPage representa a tabela channel_facebook_pages
type ChannelFacebookPage struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PageID           string    `gorm:"column:page_id;not null;uniqueIndex:idx_page_id_account_id" json:"page_id"`
	UserAccessToken  string    `gorm:"column:user_access_token;not null" json:"user_access_token"`
	PageAccessToken  string    `gorm:"column:page_access_token;not null" json:"page_access_token"`
	AccountID        uint      `gorm:"column:account_id;not null;uniqueIndex:idx_page_id_account_id" json:"account_id"`
	CreatedAt        time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	InstagramID      string    `gorm:"column:instagram_id" json:"instagram_id"`
}

// TableName define o nome da tabela para o GORM
func (ChannelFacebookPage) TableName() string {
	return "channel_facebook_pages"
}

// CreateIndexes cria os índices para a tabela ChannelFacebookPage
func (ChannelFacebookPage) CreateIndexes(db *gorm.DB) {
	// Criar índices únicos
	db.Migrator().CreateIndex(&ChannelFacebookPage{}, "index_channel_facebook_pages_on_page_id_and_account_id")
	db.Migrator().CreateIndex(&ChannelFacebookPage{}, "index_channel_facebook_pages_on_page_id")
}
