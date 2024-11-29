package models

import (
	"time"

	"gorm.io/gorm"
)

// CustomRole representa a tabela custom_roles
type CustomRole struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	AccountID   uint      `gorm:"column:account_id;not null" json:"account_id"`
	Permissions []string  `gorm:"column:permissions;type:text[];default:'{}'" json:"permissions"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (CustomRole) TableName() string {
	return "custom_roles"
}

// CreateIndexes cria os índices para a tabela CustomRole
func (CustomRole) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&CustomRole{}, "index_custom_roles_on_account_id")
}
