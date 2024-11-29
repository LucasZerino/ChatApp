package models

import (
	"time"
)

// Account representa a tabela accounts
type Account struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id"`                        // ID autoincremental
	Name             string         `gorm:"column:name;not null" json:"name"`                         // Nome (obrigatório)
	CreatedAt        time.Time      `gorm:"column:created_at;not null" json:"created_at"`             // Timestamp de criação (obrigatório)
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`             // Timestamp de atualização (obrigatório)
	Locale           int            `gorm:"column:locale;default:0" json:"locale"`                   // Valor padrão: 0
	Domain           *string        `gorm:"column:domain;size:100" json:"domain"`                    // Limite de 100 caracteres
	SupportEmail     *string        `gorm:"column:support_email;size:100" json:"support_email"`      // Limite de 100 caracteres
	FeatureFlags     int64          `gorm:"column:feature_flags;default:0;not null" json:"feature_flags"` // Valor padrão: 0, não nulo
	AutoResolveDuration *int        `gorm:"column:auto_resolve_duration" json:"auto_resolve_duration"` // Duração opcional
	Limits           JSONB          `gorm:"column:limits;default:'{}'" json:"limits"`                // Campo JSONB com valor padrão
	CustomAttributes JSONB          `gorm:"column:custom_attributes;default:'{}'" json:"custom_attributes"` // Campo JSONB com valor padrão
	Status           int            `gorm:"column:status;default:0" json:"status"`                   // Valor padrão: 0
}

// TableName define o nome da tabela para o GORM
func (Account) TableName() string {
	return "accounts"
}
