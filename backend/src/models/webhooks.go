package models

import (
	"gorm.io/gorm"
)

type Webhook struct {
	gorm.Model
	AccountID   int      `gorm:"index;not null"`
	InboxID     int      `gorm:"not null"`
	URL         string   `gorm:"type:varchar(255);uniqueIndex:idx_account_url;not null"`
	WebhookType int      `gorm:"default:0"`
	Subscriptions []string `gorm:"type:jsonb;default:'[\"conversation_status_changed\", \"conversation_updated\", \"conversation_created\", \"contact_created\", \"contact_updated\", \"message_created\", \"message_updated\", \"webwidget_triggered\"]'"`
}

// Migrate the schema
func (Webhook) TableName() string {
	return "webhooks"
}
