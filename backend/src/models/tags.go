package models

import (
	"time"
)

// Tag representa a tabela tags
type Tag struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"column:name;not null;unique" json:"name"`
	TaggingsCount  int       `gorm:"column:taggings_count;default:0" json:"taggings_count"`
	CreatedAt      time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (Tag) TableName() string {
	return "tags"
}
