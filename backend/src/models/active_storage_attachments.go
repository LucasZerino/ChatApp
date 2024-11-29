package models

import (
	"time"

	"gorm.io/gorm"
)

// ActiveStorageAttachment representa a tabela active_storage_attachments
type ActiveStorageAttachment struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`                      // Chave primária
	Name       string    `gorm:"column:name;not null" json:"name"`                         // Nome do anexo
	RecordType string    `gorm:"column:record_type;not null" json:"record_type"`           // Tipo de registro
	RecordID   uint      `gorm:"column:record_id;not null" json:"record_id"`               // ID do registro
	BlobID     uint      `gorm:"column:blob_id;not null" json:"blob_id"`                   // ID do blob
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`             // Timestamp de criação
}

// TableName define o nome da tabela para o GORM
func (ActiveStorageAttachment) TableName() string {
	return "active_storage_attachments"
}

// Após a criação, definimos os índices compostos manualmente usando migrações
func (ActiveStorageAttachment) AfterCreate(tx *gorm.DB) error {
	if err := tx.Exec(`
		CREATE INDEX IF NOT EXISTS index_active_storage_attachments_on_blob_id ON active_storage_attachments(blob_id);
		CREATE UNIQUE INDEX IF NOT EXISTS index_active_storage_attachments_uniqueness ON active_storage_attachments(record_type, record_id, name, blob_id);
	`).Error; err != nil {
		return err
	}
	return nil
}
