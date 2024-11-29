package models

import (
	"time"
)

// NotificationSubscription representa a tabela de assinaturas de notificações
type NotificationSubscription struct {
	ID                    uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID                uint           `gorm:"column:user_id;not null" json:"user_id"`
	SubscriptionType      int            `gorm:"column:subscription_type;not null" json:"subscription_type"`
	SubscriptionAttributes JSONB         `gorm:"column:subscription_attributes;default:'{}';not null" json:"subscription_attributes"`
	CreatedAt             time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	Identifier            string         `gorm:"column:identifier" json:"identifier"`
}

// TableName define o nome da tabela para o GORM
func (NotificationSubscription) TableName() string {
	return "notification_subscriptions"
}
