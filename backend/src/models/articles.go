package models

import (
	"time"

	"gorm.io/gorm"
)

// Article representa a tabela articles
type Article struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID         int       `gorm:"column:account_id;not null" json:"account_id"`
	PortalID          int       `gorm:"column:portal_id;not null" json:"portal_id"`
	CategoryID        *int      `gorm:"column:category_id" json:"category_id"`
	FolderID          *int      `gorm:"column:folder_id" json:"folder_id"`
	Title             string    `gorm:"column:title" json:"title"`
	Description       string    `gorm:"column:description" json:"description"`
	Content           string    `gorm:"column:content" json:"content"`
	Status            *int      `gorm:"column:status" json:"status"`
	Views             *int      `gorm:"column:views" json:"views"`
	CreatedAt         time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	AuthorID          *uint     `gorm:"column:author_id" json:"author_id"`
	AssociatedArticleID *uint   `gorm:"column:associated_article_id" json:"associated_article_id"`
	Meta              map[string]interface{} `gorm:"column:meta;default:'{}'" json:"meta"`
	Slug              string    `gorm:"column:slug;not null" json:"slug"`
	Position          *int      `gorm:"column:position" json:"position"`
}

// TableName define o nome da tabela para o GORM
func (Article) TableName() string {
	return "articles"
}

// Migrations para índices
func (Article) Migrate(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&Article{})
	if err != nil {
		return err
	}

	// Índices
	db.Migrator().CreateIndex(&Article{}, "index_articles_on_associated_article_id")
	db.Migrator().CreateIndex(&Article{}, "index_articles_on_author_id")
	db.Migrator().CreateIndex(&Article{}, "index_articles_on_slug")

	return nil
}
