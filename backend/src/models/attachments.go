package models

import (
	"time"

	"gorm.io/gorm"
)

// Attachment representa a tabela attachments
type Attachment struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FileType          int       `gorm:"column:file_type;default:0" json:"file_type"`
	ExternalURL       *string   `gorm:"column:external_url" json:"external_url"`
	CoordinatesLat    float64   `gorm:"column:coordinates_lat;default:0.0" json:"coordinates_lat"`
	CoordinatesLong   float64   `gorm:"column:coordinates_long;default:0.0" json:"coordinates_long"`
	MessageID         int       `gorm:"column:message_id;not null" json:"message_id"`
	AccountID         int       `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	FallbackTitle     *string   `gorm:"column:fallback_title" json:"fallback_title"`
	Extension         *string   `gorm:"column:extension" json:"extension"`
}

// TableName define o nome da tabela para o GORM
func (Attachment) TableName() string {
	return "attachments"
}

// Migrations para índices
func (Attachment) Migrate(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&Attachment{})
	if err != nil {
		return err
	}

	// Índices
	db.Migrator().CreateIndex(&Attachment{}, "index_attachments_on_account_id")
	db.Migrator().CreateIndex(&Attachment{}, "index_attachments_on_message_id")

	return nil
}
