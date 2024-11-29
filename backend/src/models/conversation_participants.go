package models

import (
	"time"

	"gorm.io/gorm"
)

// ConversationParticipant representa a tabela conversation_participants
type ConversationParticipant struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID     uint      `gorm:"column:account_id;not null" json:"account_id"`
	UserID        uint      `gorm:"column:user_id;not null" json:"user_id"`
	ConversationID uint     `gorm:"column:conversation_id;not null" json:"conversation_id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (ConversationParticipant) TableName() string {
	return "conversation_participants"
}

// CreateIndexes cria os índices para a tabela ConversationParticipant
func (ConversationParticipant) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&ConversationParticipant{}, "index_conversation_participants_on_account_id")
	db.Migrator().CreateIndex(&ConversationParticipant{}, "index_conversation_participants_on_conversation_id")
	db.Migrator().CreateIndex(&ConversationParticipant{}, "index_conversation_participants_on_user_id_and_conversation_id")
	db.Migrator().CreateIndex(&ConversationParticipant{}, "index_conversation_participants_on_user_id")
}
