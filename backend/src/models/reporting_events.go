package models

import (
	"time"
)

// ReportingEvent representa a tabela reporting_events
type ReportingEvent struct {
	ID                   uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                 string    `gorm:"column:name;not null" json:"name"`
	Value                float64   `gorm:"column:value" json:"value"`
	AccountID            uint      `gorm:"column:account_id;not null" json:"account_id"`
	InboxID              uint      `gorm:"column:inbox_id" json:"inbox_id"`
	UserID               uint      `gorm:"column:user_id" json:"user_id"`
	ConversationID       uint      `gorm:"column:conversation_id" json:"conversation_id"`
	CreatedAt            time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	ValueInBusinessHours float64   `gorm:"column:value_in_business_hours" json:"value_in_business_hours"`
	EventStartTime       time.Time `gorm:"column:event_start_time" json:"event_start_time"`
	EventEndTime         time.Time `gorm:"column:event_end_time" json:"event_end_time"`
}

// TableName define o nome da tabela para o GORM
func (ReportingEvent) TableName() string {
	return "reporting_events"
}
