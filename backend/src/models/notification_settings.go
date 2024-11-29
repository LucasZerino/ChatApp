package models

import (
	"time"
)

// NotificationSetting representa a tabela de configurações de notificações
type NotificationSetting struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID  uint      `gorm:"column:account_id" json:"account_id"`
	UserID     uint      `gorm:"column:user_id" json:"user_id"`
	EmailFlags int       `gorm:"column:email_flags;default:0;not null" json:"email_flags"`
	PushFlags  int       `gorm:"column:push_flags;default:0;not null" json:"push_flags"`
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (NotificationSetting) TableName() string {
	return "notification_settings"
}
