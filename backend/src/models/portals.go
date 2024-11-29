package models

import (
	"time"
)

// Portal representa a tabela portals
type Portal struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID         uint      `gorm:"column:account_id;not null" json:"account_id"`
	Name              string    `gorm:"column:name;not null" json:"name"`
	Slug              string    `gorm:"column:slug;not null;unique" json:"slug"`
	CustomDomain      string    `gorm:"column:custom_domain" json:"custom_domain"`
	Color             string    `gorm:"column:color" json:"color"`
	HomepageLink      string    `gorm:"column:homepage_link" json:"homepage_link"`
	PageTitle         string    `gorm:"column:page_title" json:"page_title"`
	HeaderText        string    `gorm:"column:header_text" json:"header_text"`
	CreatedAt         time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Config            JSONB     `gorm:"column:config;default:'{\"allowed_locales\": [\"en\"]}'" json:"config"`
	Archived          bool      `gorm:"column:archived;default:false" json:"archived"`
	ChannelWebWidgetID *uint    `gorm:"column:channel_web_widget_id" json:"channel_web_widget_id"`
}

// TableName define o nome da tabela para o GORM
func (Portal) TableName() string {
	return "portals"
}
