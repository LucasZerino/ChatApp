package models

import (
	"time"
)

// ActionMailboxInboundEmail representa a tabela action_mailbox_inbound_emails
type ActionMailboxInboundEmail struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`                            // Chave primária
	Status          int       `gorm:"column:status;default:0;not null" json:"status"`                // Status com valor padrão 0 e obrigatório
	MessageID       string    `gorm:"column:message_id;not null;uniqueIndex:idx_message" json:"message_id"` // Campo obrigatório com índice único
	MessageChecksum string    `gorm:"column:message_checksum;not null;uniqueIndex:idx_message" json:"message_checksum"` // Campo obrigatório com índice único
	CreatedAt       time.Time `gorm:"column:created_at;not null" json:"created_at"`                  // Timestamp de criação obrigatório
	UpdatedAt       time.Time `gorm:"column:updated_at;not null" json:"updated_at"`                  // Timestamp de atualização obrigatório
}

// TableName define o nome da tabela para o GORM
func (ActionMailboxInboundEmail) TableName() string {
	return "action_mailbox_inbound_emails"
}
