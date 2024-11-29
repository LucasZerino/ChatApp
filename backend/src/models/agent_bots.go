package models

import (
	"time"
)

// AgentBot representa a tabela agent_bots
type AgentBot struct {
	ID          uint                   `gorm:"primaryKey;autoIncrement" json:"id"`                  // Chave primária
	Name        string                 `gorm:"column:name" json:"name"`                             // Nome do bot
	Description string                 `gorm:"column:description" json:"description"`               // Descrição do bot
	OutgoingURL string                 `gorm:"column:outgoing_url" json:"outgoing_url"`             // URL de saída do bot
	AccountID   *uint                  `gorm:"column:account_id" json:"account_id"`                 // ID da conta associada (opcional)
	BotType     int                    `gorm:"column:bot_type;default:0" json:"bot_type"`           // Tipo do bot com valor padrão 0
	BotConfig   map[string]interface{} `gorm:"column:bot_config;type:jsonb;default:'{}'" json:"bot_config"` // Configuração do bot em JSONB (valor padrão vazio)
	CreatedAt   time.Time              `gorm:"column:created_at;not null" json:"created_at"`        // Timestamp de criação obrigatório
	UpdatedAt   time.Time              `gorm:"column:updated_at;not null" json:"updated_at"`        // Timestamp de atualização obrigatório
}

// TableName define o nome da tabela para o GORM
func (AgentBot) TableName() string {
	return "agent_bots"
}
