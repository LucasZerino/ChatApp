package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelLine representa a tabela channel_line
type ChannelLine struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID         uint      `gorm:"column:account_id;not null" json:"account_id"`
	LineChannelID     string    `gorm:"column:line_channel_id;not null;uniqueIndex:idx_line_channel_id" json:"line_channel_id"`
	LineChannelSecret string    `gorm:"column:line_channel_secret;not null" json:"line_channel_secret"`
	LineChannelToken  string    `gorm:"column:line_channel_token;not null" json:"line_channel_token"`
	CreatedAt         time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (ChannelLine) TableName() string {
	return "channel_line"
}

// CreateIndexes cria os índices para a tabela ChannelLine
func (ChannelLine) CreateIndexes(db *gorm.DB) {
	// Criar índice único para o campo line_channel_id
	db.Migrator().CreateIndex(&ChannelLine{}, "index_channel_line_on_line_channel_id")
}
