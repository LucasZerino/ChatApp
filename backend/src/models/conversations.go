package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Conversation representa a tabela conversations
type Conversation struct {
	ID                        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID                 uint           `gorm:"column:account_id;not null" json:"account_id"`
	InboxID                   uint           `gorm:"column:inbox_id;not null" json:"inbox_id"`
	Status                    int            `gorm:"column:status;default:0;not null" json:"status"`
	AssigneeID                *uint          `gorm:"column:assignee_id" json:"assignee_id"`
	CreatedAt                 time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                 time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	ContactID                 *uint          `gorm:"column:contact_id" json:"contact_id"`
	DisplayID                 uint           `gorm:"column:display_id;not null" json:"display_id"`
	ContactLastSeenAt         *time.Time     `gorm:"column:contact_last_seen_at" json:"contact_last_seen_at"`
	AgentLastSeenAt           *time.Time     `gorm:"column:agent_last_seen_at" json:"agent_last_seen_at"`
	AdditionalAttributes      json.RawMessage `gorm:"column:additional_attributes;default:'{}'" json:"additional_attributes"`
	ContactInboxID            *uint          `gorm:"column:contact_inbox_id" json:"contact_inbox_id"`
	UUID                      string         `gorm:"column:uuid;default:gen_random_uuid();not null" json:"uuid"`
	Identifier                *string        `gorm:"column:identifier" json:"identifier"`
	LastActivityAt            time.Time      `gorm:"column:last_activity_at;default:CURRENT_TIMESTAMP;not null" json:"last_activity_at"`
	TeamID                    *uint          `gorm:"column:team_id" json:"team_id"`
	CampaignID                *uint          `gorm:"column:campaign_id" json:"campaign_id"`
	SnoozedUntil              *time.Time     `gorm:"column:snoozed_until" json:"snoozed_until"`
	CustomAttributes          json.RawMessage `gorm:"column:custom_attributes;default:'{}'" json:"custom_attributes"`
	AssigneeLastSeenAt        *time.Time     `gorm:"column:assignee_last_seen_at" json:"assignee_last_seen_at"`
	FirstReplyCreatedAt       *time.Time     `gorm:"column:first_reply_created_at" json:"first_reply_created_at"`
	Priority                  *int           `gorm:"column:priority" json:"priority"`
	SLAPolicyID               *uint          `gorm:"column:sla_policy_id" json:"sla_policy_id"`
	WaitingSince              *time.Time     `gorm:"column:waiting_since" json:"waiting_since"`
	CachedLabelList           *string        `gorm:"column:cached_label_list" json:"cached_label_list"`
}

// TableName define o nome da tabela para o GORM
func (Conversation) TableName() string {
	return "conversations"
}

// CreateIndexes cria os índices para a tabela Conversation
func (Conversation) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_account_id_and_display_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_id_and_account_id")
	db.Migrator().CreateIndex(&Conversation{}, "conv_acid_inbid_stat_asgnid_idx")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_account_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_assignee_id_and_account_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_campaign_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_contact_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_contact_inbox_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_first_reply_created_at")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_inbox_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_priority")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_status_and_account_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_status_and_priority")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_team_id")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_uuid")
	db.Migrator().CreateIndex(&Conversation{}, "index_conversations_on_waiting_since")
}
