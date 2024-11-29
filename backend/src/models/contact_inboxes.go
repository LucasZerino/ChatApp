package models

import (
	"time"

	"gorm.io/gorm"
)

// ContactInbox representa a tabela contact_inboxes
type ContactInbox struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ContactID      uint      `gorm:"column:contact_id" json:"contact_id"`
	InboxID        uint      `gorm:"column:inbox_id" json:"inbox_id"`
	SourceID       string    `gorm:"column:source_id;not null" json:"source_id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	HmacVerified   bool      `gorm:"column:hmac_verified;default:false" json:"hmac_verified"`
	PubsubToken    string    `gorm:"column:pubsub_token" json:"pubsub_token"`
	Contact        Contact   `gorm:"foreignKey:ContactID" json:"contact"` // Referência ao modelo Contact
}

// TableName define o nome da tabela para o GORM
func (ContactInbox) TableName() string {
	return "contact_inboxes"
}

// CreateIndexes cria os índices para a tabela ContactInbox
func (ContactInbox) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&ContactInbox{}, "index_contact_inboxes_on_contact_id")
	db.Migrator().CreateIndex(&ContactInbox{}, "index_contact_inboxes_on_inbox_id_and_source_id")
	db.Migrator().CreateIndex(&ContactInbox{}, "index_contact_inboxes_on_inbox_id")
	db.Migrator().CreateIndex(&ContactInbox{}, "index_contact_inboxes_on_pubsub_token")
	db.Migrator().CreateIndex(&ContactInbox{}, "index_contact_inboxes_on_source_id")
}
