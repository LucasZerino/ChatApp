package models

import (
	"time"
)

// SlaEvent representa a tabela sla_events
type SlaEvent struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AppliedSlaID    uint      `gorm:"column:applied_sla_id;not null" json:"applied_sla_id"`
	ConversationID  uint      `gorm:"column:conversation_id;not null" json:"conversation_id"`
	AccountID       uint      `gorm:"column:account_id;not null" json:"account_id"`
	SlaPolicyID     uint      `gorm:"column:sla_policy_id;not null" json:"sla_policy_id"`
	InboxID         uint      `gorm:"column:inbox_id;not null" json:"inbox_id"`
	EventType       int       `gorm:"column:event_type" json:"event_type"`
	Meta            JSONB     `gorm:"column:meta;default:'{}'" json:"meta"`
	CreatedAt       time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (SlaEvent) TableName() string {
	return "sla_events"
}
