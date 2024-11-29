package models

import (
	"time"
)

// InstallationConfig representa a tabela installation_configs
type InstallationConfig struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string         `gorm:"column:name;not null;unique" json:"name"`
	SerializedValue string         `gorm:"column:serialized_value;type:jsonb;default:'{}';not null" json:"serialized_value"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	Locked          bool           `gorm:"column:locked;default:true;not null" json:"locked"`
}

// TableName define o nome da tabela para o GORM
func (InstallationConfig) TableName() string {
	return "installation_configs"
}
