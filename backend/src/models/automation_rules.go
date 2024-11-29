package models

import (
	"time"

	"gorm.io/gorm"
)

// AutomationRule representa a tabela automation_rules
type AutomationRule struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   uint      `gorm:"column:account_id;not null" json:"account_id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	EventName   string    `gorm:"column:event_name;not null" json:"event_name"`
	Conditions  string    `gorm:"column:conditions;type:jsonb;default:'{}';not null" json:"conditions"`
	Actions     string    `gorm:"column:actions;type:jsonb;default:'{}';not null" json:"actions"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Active      bool      `gorm:"column:active;default:true;not null" json:"active"`
}

// TableName define o nome da tabela para o GORM
func (AutomationRule) TableName() string {
	return "automation_rules"
}

// Migrations para índices
func (AutomationRule) Migrate(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&AutomationRule{})
	if err != nil {
		return err
	}

	// Índices
	db.Migrator().CreateIndex(&AutomationRule{}, "index_automation_rules_on_account_id")

	return nil
}
