package models

import (
	"time"
)

// Label representa a tabela labels
type Label struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	Color       string     `gorm:"column:color;default:'#1f93ff';not null" json:"color"`
	ShowOnSidebar *bool    `gorm:"column:show_on_sidebar" json:"show_on_sidebar"`
	AccountID   *int       `gorm:"column:account_id" json:"account_id"`
	CreatedAt   time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Label) TableName() string {
	return "labels"
}
