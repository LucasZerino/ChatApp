package models

import (
	"time"
)

// Mention representa a tabela mentions
type Mention struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint       `gorm:"column:user_id;not null" json:"user_id"`
	ConversationID uint      `gorm:"column:conversation_id;not null" json:"conversation_id"`
	AccountID     uint       `gorm:"column:account_id;not null" json:"account_id"`
	MentionedAt   time.Time  `gorm:"column:mentioned_at;not null" json:"mentioned_at"`
	CreatedAt     time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Mention) TableName() string {
	return "mentions"
}
