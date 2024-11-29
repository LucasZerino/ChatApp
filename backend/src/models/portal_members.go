package models

import (
	"time"
)

// PortalMember representa a tabela portal_members
type PortalMember struct {
	PortalID  uint      `gorm:"column:portal_id;not null" json:"portal_id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (PortalMember) TableName() string {
	return "portal_members"
}
