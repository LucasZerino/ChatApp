package models

import (
	"time"
)

// SlaPolicy representa a tabela sla_policies
type SlaPolicy struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                      string    `gorm:"column:name;not null" json:"name"`
	FirstResponseTimeThreshold *float64 `gorm:"column:first_response_time_threshold" json:"first_response_time_threshold"`
	NextResponseTimeThreshold  *float64 `gorm:"column:next_response_time_threshold" json:"next_response_time_threshold"`
	OnlyDuringBusinessHours   bool      `gorm:"column:only_during_business_hours;default:false" json:"only_during_business_hours"`
	AccountID                 uint      `gorm:"column:account_id;not null" json:"account_id"`
	Description               *string   `gorm:"column:description" json:"description"`
	ResolutionTimeThreshold   *float64 `gorm:"column:resolution_time_threshold" json:"resolution_time_threshold"`
	CreatedAt                 time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (SlaPolicy) TableName() string {
	return "sla_policies"
}
