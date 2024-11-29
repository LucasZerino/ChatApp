package models

import (
	"time"
)

// ActiveStorageBlob representa a tabela active_storage_blobs
type ActiveStorageBlob struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`             // Chave primária
	Key         string    `gorm:"column:key;not null;uniqueIndex" json:"key"`     // Chave única
	Filename    string    `gorm:"column:filename;not null" json:"filename"`       // Nome do arquivo
	ContentType *string   `gorm:"column:content_type" json:"content_type"`        // Tipo de conteúdo (opcional)
	Metadata    *string   `gorm:"column:metadata" json:"metadata"`                // Metadados (opcional)
	ByteSize    int64     `gorm:"column:byte_size;not null" json:"byte_size"`     // Tamanho em bytes
	Checksum    *string   `gorm:"column:checksum" json:"checksum"`                // Checksum (opcional)
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`   // Timestamp de criação
	ServiceName string    `gorm:"column:service_name;not null" json:"service_name"` // Nome do serviço
}

// TableName define o nome da tabela para o GORM
func (ActiveStorageBlob) TableName() string {
	return "active_storage_blobs"
}
