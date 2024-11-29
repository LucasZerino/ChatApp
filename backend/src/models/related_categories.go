package models

import (
	"time"
)

// RelatedCategory representa a tabela related_categories
type RelatedCategory struct {
	CategoryID        uint      `gorm:"column:category_id;not null" json:"category_id"`
	RelatedCategoryID uint      `gorm:"column:related_category_id;not null" json:"related_category_id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (RelatedCategory) TableName() string {
	return "related_categories"
}
