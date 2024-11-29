package models

import (
	"time"

	"gorm.io/gorm"
)

// Contact representa a tabela contacts
type Contact struct {
	ID                    uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                  string    `gorm:"column:name;default:''" json:"name"`
	Email                 string    `gorm:"column:email" json:"email"`
	PhoneNumber           string    `gorm:"column:phone_number" json:"phone_number"`
	AccountID             uint      `gorm:"column:account_id;not null" json:"account_id"`
	CreatedAt             time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	AdditionalAttributes  string    `gorm:"column:additional_attributes;type:jsonb;default:'{}'" json:"additional_attributes"`
	Identifier            string    `gorm:"column:identifier" json:"identifier"`
	CustomAttributes      string    `gorm:"column:custom_attributes;type:jsonb;default:'{}'" json:"custom_attributes"`
	LastActivityAt        time.Time `gorm:"column:last_activity_at" json:"last_activity_at"`
	ContactType           int       `gorm:"column:contact_type;default:0" json:"contact_type"`
	MiddleName            string    `gorm:"column:middle_name;default:''" json:"middle_name"`
	LastName              string    `gorm:"column:last_name;default:''" json:"last_name"`
	Location              string    `gorm:"column:location;default:''" json:"location"`
	CountryCode           string    `gorm:"column:country_code;default:''" json:"country_code"`
	Blocked               bool      `gorm:"column:blocked;default:false;not null" json:"blocked"`
}

// TableName define o nome da tabela para o GORM
func (Contact) TableName() string {
	return "contacts"
}

// CreateIndexes cria os índices para a tabela Contact
func (Contact) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_lower_email_account_id")
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_nonempty_fields")
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_account_id_and_last_activity_at")
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_account_id")
	db.Migrator().CreateIndex(&Contact{}, "index_resolved_contact_account_id")
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_blocked")
	db.Migrator().CreateIndex(&Contact{}, "uniq_email_per_account_contact")
	db.Migrator().CreateIndex(&Contact{}, "uniq_identifier_per_account_contact")
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_name_email_phone_number_identifier")
	db.Migrator().CreateIndex(&Contact{}, "index_contacts_on_phone_number_and_account_id")
}
