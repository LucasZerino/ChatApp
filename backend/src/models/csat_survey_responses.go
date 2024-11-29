package models

import (
	"time"

	"gorm.io/gorm"
)

// CsatSurveyResponse representa a tabela csat_survey_responses
type CsatSurveyResponse struct {
	ID                  uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID           uint           `gorm:"column:account_id;not null" json:"account_id"`
	ConversationID      uint           `gorm:"column:conversation_id;not null" json:"conversation_id"`
	MessageID           uint           `gorm:"column:message_id;not null" json:"message_id"`
	Rating              int            `gorm:"column:rating;not null" json:"rating"`
	FeedbackMessage     *string        `gorm:"column:feedback_message" json:"feedback_message"`
	ContactID           uint           `gorm:"column:contact_id;not null" json:"contact_id"`
	AssignedAgentID     *uint          `gorm:"column:assigned_agent_id" json:"assigned_agent_id"`
	CreatedAt           time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (CsatSurveyResponse) TableName() string {
	return "csat_survey_responses"
}

// CreateIndexes cria os índices para a tabela CsatSurveyResponse
func (CsatSurveyResponse) CreateIndexes(db *gorm.DB) {
	// Criar índices para os campos
	db.Migrator().CreateIndex(&CsatSurveyResponse{}, "index_csat_survey_responses_on_account_id")
	db.Migrator().CreateIndex(&CsatSurveyResponse{}, "index_csat_survey_responses_on_assigned_agent_id")
	db.Migrator().CreateIndex(&CsatSurveyResponse{}, "index_csat_survey_responses_on_contact_id")
	db.Migrator().CreateIndex(&CsatSurveyResponse{}, "index_csat_survey_responses_on_conversation_id")
	db.Migrator().CreateIndex(&CsatSurveyResponse{}, "index_csat_survey_responses_on_message_id")
}
