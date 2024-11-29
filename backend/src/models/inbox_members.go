package models

import (
	"time"
)

// InboxMember representa a tabela inbox_members
type InboxMember struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	InboxID   uint      `gorm:"column:inbox_id;not null" json:"inbox_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (InboxMember) TableName() string {
	return "inbox_members"
}
