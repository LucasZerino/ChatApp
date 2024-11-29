package models

import (
	"time"
)

// Team representa a tabela teams
type Team struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Description    string    `gorm:"column:description" json:"description"`
	AllowAutoAssign bool     `gorm:"column:allow_auto_assign;default:true" json:"allow_auto_assign"`
	AccountID      uint      `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Team) TableName() string {
	return "teams"
}
