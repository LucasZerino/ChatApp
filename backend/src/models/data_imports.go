package models

import (
	"time"

	"gorm.io/gorm"
)

// DataImport representa a tabela data_imports
type DataImport struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID       uint      `gorm:"column:account_id;not null" json:"account_id"`
	DataType        string    `gorm:"column:data_type;not null" json:"data_type"`
	Status          int       `gorm:"column:status;default:0;not null" json:"status"`
	ProcessingErrors string   `gorm:"column:processing_errors" json:"processing_errors,omitempty"`
	TotalRecords    int       `gorm:"column:total_records" json:"total_records,omitempty"`
	ProcessedRecords int      `gorm:"column:processed_records" json:"processed_records,omitempty"`
	CreatedAt       time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (DataImport) TableName() string {
	return "data_imports"
}

// CreateIndexes cria os índices para a tabela DataImport
func (DataImport) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&DataImport{}, "index_data_imports_on_account_id")
}
