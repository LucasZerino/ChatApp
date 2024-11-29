package models

import (
	"time"
)

// Tagging representa a tabela taggings
type Tagging struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TagID         uint      `gorm:"column:tag_id;not null" json:"tag_id"`
	TaggableType  string    `gorm:"column:taggable_type;not null" json:"taggable_type"`
	TaggableID    uint      `gorm:"column:taggable_id;not null" json:"taggable_id"`
	TaggerType    string    `gorm:"column:tagger_type" json:"tagger_type"`
	TaggerID      uint      `gorm:"column:tagger_id" json:"tagger_id"`
	Context       *string   `gorm:"column:context;size:128" json:"context"` // Context é opcional, então é um ponteiro
	CreatedAt     time.Time `gorm:"column:created_at;not null" json:"created_at"`
}

// TableName define o nome da tabela para o GORM
func (Tagging) TableName() string {
	return "taggings"
}
