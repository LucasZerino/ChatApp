package models

import (
	"time"
)

// AccountUser representa a tabela account_users
type AccountUser struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`                                // Chave primária (opcional, se não for usado no Ruby, remova)
	AccountID     int64      `gorm:"column:account_id;index" json:"account_id"`                         // Índice simples
	UserID        int64      `gorm:"column:user_id;index" json:"user_id"`                               // Índice simples
	Role          int        `gorm:"column:role;default:0" json:"role"`                                 // Valor padrão: 0
	InviterID     *int64     `gorm:"column:inviter_id" json:"inviter_id"`                               // Pode ser nulo
	CreatedAt     time.Time  `gorm:"column:created_at;not null" json:"created_at"`                      // Timestamp de criação
	UpdatedAt     time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`                      // Timestamp de atualização
	ActiveAt      *time.Time `gorm:"column:active_at" json:"active_at"`                                 // Timestamp opcional
	Availability  int        `gorm:"column:availability;default:0;not null" json:"availability"`        // Valor padrão: 0, não nulo
	AutoOffline   bool       `gorm:"column:auto_offline;default:true;not null" json:"auto_offline"`     // Valor padrão: true, não nulo
	CustomRoleID  *int64     `gorm:"column:custom_role_id;index" json:"custom_role_id"`                 // Índice simples
}

// TableName define o nome da tabela para o GORM
func (AccountUser) TableName() string {
	return "account_users"
}
