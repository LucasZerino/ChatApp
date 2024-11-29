package models

import "gorm.io/gorm"

// ActiveStorageVariantRecord representa a tabela active_storage_variant_records
type ActiveStorageVariantRecord struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`                           // Chave primária
	BlobID          uint   `gorm:"column:blob_id;not null" json:"blob_id"`                      // ID do blob, obrigatório
	VariationDigest string `gorm:"column:variation_digest;not null" json:"variation_digest"`     // Digest da variação, obrigatório
}

// TableName define o nome da tabela para o GORM
func (ActiveStorageVariantRecord) TableName() string {
	return "active_storage_variant_records"
}

// Índices para a tabela
func (ActiveStorageVariantRecord) AfterCreate(tx *gorm.DB) error {
	if err := tx.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS index_active_storage_variant_records_uniqueness
		ON active_storage_variant_records (blob_id, variation_digest);
	`).Error; err != nil {
		return err
	}
	return nil
}
