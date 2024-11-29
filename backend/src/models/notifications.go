package models

import (
	"time"
)

// Notification representa a tabela de notificações
type Notification struct {
	ID                  uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID           uint           `gorm:"column:account_id;not null" json:"account_id"`
	UserID              uint           `gorm:"column:user_id;not null" json:"user_id"`
	NotificationType    int            `gorm:"column:notification_type;not null" json:"notification_type"`
	PrimaryActorType    string         `gorm:"column:primary_actor_type;not null" json:"primary_actor_type"`
	PrimaryActorID      uint           `gorm:"column:primary_actor_id;not null" json:"primary_actor_id"`
	SecondaryActorType  *string        `gorm:"column:secondary_actor_type" json:"secondary_actor_type"`
	SecondaryActorID    *uint          `gorm:"column:secondary_actor_id" json:"secondary_actor_id"`
	ReadAt              *time.Time     `gorm:"column:read_at" json:"read_at"`
	CreatedAt           time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	SnoozedUntil        *time.Time     `gorm:"column:snoozed_until" json:"snoozed_until"`
	LastActivityAt      time.Time      `gorm:"column:last_activity_at;default:CURRENT_TIMESTAMP" json:"last_activity_at"`
	Meta                JSONB          `gorm:"column:meta;default:'{}'" json:"meta"`
}

// TableName define o nome da tabela para o GORM
func (Notification) TableName() string {
	return "notifications"
}
