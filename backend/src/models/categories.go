package models

import (
	"time"

	"gorm.io/gorm"
)

// Category representa a tabela categories
type Category struct {
	ID                    uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID             uint      `gorm:"column:account_id;not null;index" json:"account_id"`
	PortalID              uint      `gorm:"column:portal_id;not null" json:"portal_id"`
	Name                  string    `gorm:"column:name" json:"name"`
	Description           string    `gorm:"column:description" json:"description"`
	Position              int       `gorm:"column:position" json:"position"`
	CreatedAt            time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Locale               string    `gorm:"column:locale;default:'en';index" json:"locale"`
	Slug                 string    `gorm:"column:slug;not null;index:slug_locale_portal_id" json:"slug"`
	ParentCategoryID     *uint     `gorm:"column:parent_category_id;index" json:"parent_category_id"`
	AssociatedCategoryID *uint     `gorm:"column:associated_category_id;index" json:"associated_category_id"`
	Icon                 string    `gorm:"column:icon;default:''" json:"icon"`
}

// TableName define o nome da tabela para o GORM
func (Category) TableName() string {
	return "categories"
}

// Migrar e criar índices manualmente
func (Category) CreateIndexes(db *gorm.DB) {
	// Criar índices compostos ou únicos
	db.Migrator().CreateIndex(&Category{}, "index_categories_on_associated_category_id")
	db.Migrator().CreateIndex(&Category{}, "index_categories_on_locale_and_account_id")
	db.Migrator().CreateIndex(&Category{}, "index_categories_on_locale")
	db.Migrator().CreateIndex(&Category{}, "index_categories_on_parent_category_id")
	db.Migrator().CreateIndex(&Category{}, "index_categories_on_slug_and_locale_and_portal_id")
}
