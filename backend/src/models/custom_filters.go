package models

import (
	"time"

	"gorm.io/gorm"
)

// CustomFilter representa a tabela custom_filters
type CustomFilter struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	FilterType int           `gorm:"column:filter_type;default:0;not null" json:"filter_type"`
	Query     []byte         `gorm:"column:query;type:jsonb;default:'{}';not null" json:"query"`
	AccountID uint           `gorm:"column:account_id;not null" json:"account_id"`
	UserID    uint           `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (CustomFilter) TableName() string {
	return "custom_filters"
}

// CreateIndexes cria os índices para a tabela CustomFilter
func (CustomFilter) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&CustomFilter{}, "index_custom_filters_on_account_id")
	db.Migrator().CreateIndex(&CustomFilter{}, "index_custom_filters_on_user_id")
}
