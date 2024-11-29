package models

import (
	"time"
)

// TeamMember representa a tabela team_members
type TeamMember struct {
	TeamID     uint      `gorm:"column:team_id;not null" json:"team_id"`
	UserID     uint      `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (TeamMember) TableName() string {
	return "team_members"
}
