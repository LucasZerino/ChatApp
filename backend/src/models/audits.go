package models

import (
	"time"

	"gorm.io/gorm"
)

// Audit representa a tabela audits
type Audit struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AuditableID     uint      `gorm:"column:auditable_id" json:"auditable_id"`
	AuditableType   string    `gorm:"column:auditable_type" json:"auditable_type"`
	AssociatedID    uint      `gorm:"column:associated_id" json:"associated_id"`
	AssociatedType  string    `gorm:"column:associated_type" json:"associated_type"`
	UserID          uint      `gorm:"column:user_id" json:"user_id"`
	UserType        string    `gorm:"column:user_type" json:"user_type"`
	Username        string    `gorm:"column:username" json:"username"`
	Action          string    `gorm:"column:action" json:"action"`
	AuditedChanges  string    `gorm:"column:audited_changes" json:"audited_changes"`
	Version         int       `gorm:"column:version;default:0" json:"version"`
	Comment         string    `gorm:"column:comment" json:"comment"`
	RemoteAddress   string    `gorm:"column:remote_address" json:"remote_address"`
	RequestUUID     string    `gorm:"column:request_uuid" json:"request_uuid"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName define o nome da tabela para o GORM
func (Audit) TableName() string {
	return "audits"
}

// Migrations para índices
func (Audit) Migrate(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&Audit{})
	if err != nil {
		return err
	}

	// Índices
	db.Migrator().CreateIndex(&Audit{}, "associated_index")
	db.Migrator().CreateIndex(&Audit{}, "auditable_index")
	db.Migrator().CreateIndex(&Audit{}, "index_audits_on_created_at")
	db.Migrator().CreateIndex(&Audit{}, "index_audits_on_request_uuid")
	db.Migrator().CreateIndex(&Audit{}, "user_index")

	return nil
}
