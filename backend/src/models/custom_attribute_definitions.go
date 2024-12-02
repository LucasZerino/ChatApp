package models

import (
	"time"

	"gorm.io/gorm"
)

// CustomAttributeDefinition representa a tabela custom_attribute_definitions
type CustomAttributeDefinition struct {
	ID                   uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AttributeDisplayName  string         `gorm:"column:attribute_display_name" json:"attribute_display_name"`
	AttributeKey          string         `gorm:"column:attribute_key" json:"attribute_key"`
	AttributeDisplayType  int            `gorm:"column:attribute_display_type;default:0" json:"attribute_display_type"`
	DefaultValue          int            `gorm:"column:default_value" json:"default_value"`
	AttributeModel        int            `gorm:"column:attribute_model;default:0" json:"attribute_model"`
	AccountID             uint           `gorm:"column:account_id" json:"account_id"`
	CreatedAt             time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	AttributeDescription  string         `gorm:"column:attribute_description" json:"attribute_description"`
	AttributeValues       string         `gorm:"column:attribute_values;type:jsonb;default:'[]'" json:"attribute_values"`
	RegexPattern          string         `gorm:"column:regex_pattern" json:"regex_pattern"`
	RegexCue              string         `gorm:"column:regex_cue" json:"regex_cue"`
}

// TableName define o nome da tabela para o GORM
func (CustomAttributeDefinition) TableName() string {
	return "custom_attribute_definitions"
}

// CreateIndexes cria os índices para a tabela CustomAttributeDefinition
func (CustomAttributeDefinition) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&CustomAttributeDefinition{}, "index_custom_attribute_definitions_on_account_id")
	db.Migrator().CreateIndex(&CustomAttributeDefinition{}, "attribute_key_model_index")
}
