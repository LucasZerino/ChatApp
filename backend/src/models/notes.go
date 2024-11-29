package models

import (
	"time"
)

// Note representa a tabela de notas
type Note struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"column:content;not null" json:"content"`
	AccountID uint      `gorm:"column:account_id;not null" json:"account_id"`
	ContactID uint      `gorm:"column:contact_id;not null" json:"contact_id"`
	UserID    *uint     `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Note) TableName() string {
	return "notes"
}
