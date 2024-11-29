package models

import (
	"time"
)

// CannedResponse representa a tabela canned_responses
type CannedResponse struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID  uint      `gorm:"column:account_id;not null" json:"account_id"`
	ShortCode  string    `gorm:"column:short_code" json:"short_code"`
	Content    string    `gorm:"column:content" json:"content"`
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName define o nome da tabela para o GORM
func (CannedResponse) TableName() string {
	return "canned_responses"
}
