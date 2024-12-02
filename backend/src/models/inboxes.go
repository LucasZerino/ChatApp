package models

import (
	"time"
)

// Inbox representa a tabela inboxes
type Inbox struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID                 uint      `gorm:"column:channel_id;not null" json:"channel_id"`
	AccountID                 uint      `gorm:"column:account_id;not null" json:"account_id"`
	Name                      string    `gorm:"column:name;not null" json:"name"`
	CreatedAt                 time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	ChannelType               string    `gorm:"column:channel_type" json:"channel_type"`
	EnableAutoAssignment      bool      `gorm:"column:enable_auto_assignment;default:true" json:"enable_auto_assignment"`
	GreetingEnabled           bool      `gorm:"column:greeting_enabled;default:false" json:"greeting_enabled"`
	GreetingMessage           string    `gorm:"column:greeting_message" json:"greeting_message"`
	EmailAddress              string    `gorm:"column:email_address" json:"email_address"`
	WorkingHoursEnabled       bool      `gorm:"column:working_hours_enabled;default:false" json:"working_hours_enabled"`
	OutOfOfficeMessage        string    `gorm:"column:out_of_office_message" json:"out_of_office_message"`
	Timezone                  string    `gorm:"column:timezone;default:'UTC'" json:"timezone"`
	EnableEmailCollect        bool      `gorm:"column:enable_email_collect;default:true" json:"enable_email_collect"`
	CsatSurveyEnabled         bool      `gorm:"column:csat_survey_enabled;default:false" json:"csat_survey_enabled"`
	AllowMessagesAfterResolved bool      `gorm:"column:allow_messages_after_resolved;default:true" json:"allow_messages_after_resolved"`
	AutoAssignmentConfig      JSONB     `gorm:"column:auto_assignment_config;type:jsonb;default:'{}'" json:"auto_assignment_config"`
	LockToSingleConversation  bool      `gorm:"column:lock_to_single_conversation;default:false;not null" json:"lock_to_single_conversation"`
	PortalID                  uint      `gorm:"column:portal_id" json:"portal_id"`
	SenderNameType            int       `gorm:"column:sender_name_type;default:0;not null" json:"sender_name_type"`
	BusinessName              string    `gorm:"column:business_name" json:"business_name"`
	CsatResponseVisible       bool      `gorm:"column:csat_response_visible;default:false;not null" json:"csat_response_visible"`
	AllowAgentToDeleteMessage bool      `gorm:"column:allow_agent_to_delete_message;default:true;not null" json:"allow_agent_to_delete_message"`
}

// TableName define o nome da tabela para o GORM
func (Inbox) TableName() string {
	return "inboxes"
}
