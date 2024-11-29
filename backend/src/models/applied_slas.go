package models

import (
	"time"

	"gorm.io/gorm"
)

// AppliedSla representa a tabela applied_slas
type AppliedSla struct {
	AccountID      uint      `gorm:"column:account_id;not null" json:"account_id"`
	SlaPolicyID    uint      `gorm:"column:sla_policy_id;not null" json:"sla_policy_id"`
	ConversationID uint      `gorm:"column:conversation_id;not null" json:"conversation_id"`
	SlaStatus      int       `gorm:"column:sla_status;default:0" json:"sla_status"`
	CreatedAt      time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (AppliedSla) TableName() string {
	return "applied_slas"
}

// Migrations para índices
func (AppliedSla) Migrate(db *gorm.DB) error {
	err := db.Migrator().CreateTable(&AppliedSla{})
	if err != nil {
		return err
	}

	// Índice único composto
	if !db.Migrator().HasIndex(&AppliedSla{}, "index_applied_slas_on_account_sla_policy_conversation") {
		db.Migrator().CreateIndex(&AppliedSla{}, "index_applied_slas_on_account_sla_policy_conversation")
	}

	// Índices não únicos
	db.Migrator().CreateIndex(&AppliedSla{}, "index_applied_slas_on_account_id")
	db.Migrator().CreateIndex(&AppliedSla{}, "index_applied_slas_on_conversation_id")
	db.Migrator().CreateIndex(&AppliedSla{}, "index_applied_slas_on_sla_policy_id")

	return nil
}
