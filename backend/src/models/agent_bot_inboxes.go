package models

import (
	"time"
)

// AgentBotInbox representa a tabela agent_bot_inboxes
type AgentBotInbox struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`                 // Chave primária
	InboxID    *int      `gorm:"column:inbox_id" json:"inbox_id"`                    // ID do inbox (opcional)
	AgentBotID *int      `gorm:"column:agent_bot_id" json:"agent_bot_id"`            // ID do bot do agente (opcional)
	Status     int       `gorm:"column:status;default:0" json:"status"`              // Status com valor padrão 0
	AccountID  *int      `gorm:"column:account_id" json:"account_id"`                // ID da conta (opcional)
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`       // Timestamp de criação obrigatório
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`       // Timestamp de atualização obrigatório
}

// TableName define o nome da tabela para o GORM
func (AgentBotInbox) TableName() string {
	return "agent_bot_inboxes"
}
