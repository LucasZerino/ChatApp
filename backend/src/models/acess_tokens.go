package models

import (
	"time"
)

// AccessToken representa a tabela access_tokens
type AccessToken struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`                   // ID único do token
	OwnerType  string      `gorm:"column:owner_type;index:idx_owner" json:"owner_type"`  // Tipo do proprietário, parte do índice composto
	OwnerID    int64       `gorm:"column:owner_id;index:idx_owner" json:"owner_id"`      // ID do proprietário, parte do índice composto
	Token      string      `gorm:"column:token;uniqueIndex" json:"token"`                // Token único
	CreatedAt  time.Time   `gorm:"column:created_at;not null" json:"created_at"`         // Timestamp de criação
	UpdatedAt  time.Time   `gorm:"column:updated_at;not null" json:"updated_at"`         // Timestamp de atualização
	Owner      interface{} `gorm:"polymorphic:Owner" json:"owner"`                       // Relacionamento polimórfico
}

// TableName define o nome da tabela para o GORM
func (AccessToken) TableName() string {
	return "access_tokens"
}
