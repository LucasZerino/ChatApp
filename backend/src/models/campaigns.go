package models

import (
	"time"

	"gorm.io/gorm"
)

// Campaign representa a tabela campaigns
type Campaign struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	DisplayID                 int       `gorm:"column:display_id;not null" json:"display_id"`
	Title                     string    `gorm:"column:title;not null" json:"title"`
	Description               string    `gorm:"column:description" json:"description"`
	Message                   string    `gorm:"column:message;not null" json:"message"`
	SenderID                  *int      `gorm:"column:sender_id" json:"sender_id"`
	Enabled                   bool      `gorm:"column:enabled;default:true" json:"enabled"`
	AccountID                 uint      `gorm:"column:account_id;not null" json:"account_id"`
	InboxID                   uint      `gorm:"column:inbox_id;not null" json:"inbox_id"`
	TriggerRules              string    `gorm:"column:trigger_rules;type:jsonb;default:'{}'" json:"trigger_rules"`
	CreatedAt                 time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	CampaignType              int       `gorm:"column:campaign_type;default:0;not null" json:"campaign_type"`
	CampaignStatus            int       `gorm:"column:campaign_status;default:0;not null" json:"campaign_status"`
	Audience                  string    `gorm:"column:audience;type:jsonb;default:'[]'" json:"audience"`
	ScheduledAt               *time.Time `gorm:"column:scheduled_at" json:"scheduled_at"`
	TriggerOnlyDuringBusinessHours bool `gorm:"column:trigger_only_during_business_hours;default:false" json:"trigger_only_during_business_hours"`
}

// TableName define o nome da tabela para o GORM
func (Campaign) TableName() string {
	return "campaigns"
}

// Migrations para índices
func (Campaign) Migrate(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&Campaign{})
	if err != nil {
		return err
	}

	// Índices
	db.Migrator().CreateIndex(&Campaign{}, "index_campaigns_on_account_id")
	db.Migrator().CreateIndex(&Campaign{}, "index_campaigns_on_campaign_status")
	db.Migrator().CreateIndex(&Campaign{}, "index_campaigns_on_campaign_type")
	db.Migrator().CreateIndex(&Campaign{}, "index_campaigns_on_inbox_id")
	db.Migrator().CreateIndex(&Campaign{}, "index_campaigns_on_scheduled_at")

	return nil
}
