package models

import (
	"time"
)

// Macro representa a tabela macros
type Macro struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID    uint       `gorm:"column:account_id;not null" json:"account_id"`
	Name         string     `gorm:"column:name;not null" json:"name"`
	Visibility   int        `gorm:"column:visibility;default:0" json:"visibility"`
	CreatedByID  *uint      `gorm:"column:created_by_id" json:"created_by_id"`
	UpdatedByID  *uint      `gorm:"column:updated_by_id" json:"updated_by_id"`
	Actions      JSONB      `gorm:"column:actions;default:'{}';not null" json:"actions"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Macro) TableName() string {
	return "macros"
}
