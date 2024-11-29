package models

import (
	"time"

	"gorm.io/gorm"
)

// EmailTemplate representa a tabela email_templates
type EmailTemplate struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Body       string    `gorm:"column:body;not null" json:"body"`
	AccountID  uint      `gorm:"column:account_id" json:"account_id,omitempty"`
	TemplateType int     `gorm:"column:template_type;default:1" json:"template_type"`
	Locale     int       `gorm:"column:locale;default:0;not null" json:"locale"`
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (EmailTemplate) TableName() string {
	return "email_templates"
}

// CreateIndexes cria os índices para a tabela EmailTemplate
func (EmailTemplate) CreateIndexes(db *gorm.DB) {
	// Criar índice único para o nome e account_id
	db.Migrator().CreateIndex(&EmailTemplate{}, "index_email_templates_on_name_and_account_id")
}
