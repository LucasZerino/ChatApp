package models

import (
	"time"
)

// Message representa a tabela de mensagens
type Message struct {
	ID                   uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Content              string     `gorm:"column:content;not null" json:"content"`
	AccountID            uint       `gorm:"column:account_id;not null" json:"account_id"`
	InboxID              uint       `gorm:"column:inbox_id;not null" json:"inbox_id"`
	ConversationID       uint       `gorm:"column:conversation_id;not null" json:"conversation_id"`
	MessageType          int        `gorm:"column:message_type;not null" json:"message_type"`
	CreatedAt            time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`
	Private             bool       `gorm:"column:private;default:false;not null" json:"private"`
	Status              int        `gorm:"column:status;default:0" json:"status"`
	SourceID            string     `gorm:"column:source_id" json:"source_id"`
	ContentType         int        `gorm:"column:content_type;default:0;not null" json:"content_type"`
	ContentAttributes   JSONB      `gorm:"column:content_attributes;default:'{}';not null" json:"content_attributes"`
	SenderType          string     `gorm:"column:sender_type" json:"sender_type"`
	SenderID            *uint      `gorm:"column:sender_id" json:"sender_id"`
	ExternalSourceIDs   JSONB      `gorm:"column:external_source_ids;default:'{}'" json:"external_source_ids"`
	AdditionalAttributes JSONB     `gorm:"column:additional_attributes;default:'{}'" json:"additional_attributes"`
	ProcessedMessageContent string `gorm:"column:processed_message_content" json:"processed_message_content"`
	Sentiment           JSONB      `gorm:"column:sentiment;default:'{}'" json:"sentiment"`
	Attachments         []Attachment `gorm:"foreignKey:MessageID" json:"attachments"`
}

// TableName define o nome da tabela para o GORM
func (Message) TableName() string {
	return "messages"
}

