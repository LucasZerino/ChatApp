package models

import (
	"time"
)

// IntegrationHook representa a tabela integrations_hooks
type IntegrationHook struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Status      int            `gorm:"column:status;default:1" json:"status"`
	InboxID     *int           `gorm:"column:inbox_id" json:"inbox_id"`
	AccountID   *int           `gorm:"column:account_id" json:"account_id"`
	AppID       string         `gorm:"column:app_id" json:"app_id"`
	HookType    int            `gorm:"column:hook_type;default:0" json:"hook_type"`
	ReferenceID string         `gorm:"column:reference_id" json:"reference_id"`
	AccessToken string         `gorm:"column:access_token" json:"access_token"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	Settings    *JSONB `gorm:"column:settings;type:jsonb;default:'{}'" json:"settings"`
}

// TableName define o nome da tabela para o GORM
func (IntegrationHook) TableName() string {
	return "integrations_hooks"
}
