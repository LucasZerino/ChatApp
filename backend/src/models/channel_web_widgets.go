package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelWebWidget representa a tabela channel_web_widgets
type ChannelWebWidget struct {
	ID                         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	WebsiteURL                 string    `gorm:"column:website_url" json:"website_url"`
	AccountID                  uint      `gorm:"column:account_id" json:"account_id"`
	CreatedAt                  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt                  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	WebsiteToken               string    `gorm:"column:website_token" json:"website_token"`
	WidgetColor                string    `gorm:"column:widget_color;default:'#1f93ff'" json:"widget_color"`
	WelcomeTitle              string    `gorm:"column:welcome_title" json:"welcome_title"`
	WelcomeTagline            string    `gorm:"column:welcome_tagline" json:"welcome_tagline"`
	FeatureFlags              int       `gorm:"column:feature_flags;default:7;not null" json:"feature_flags"`
	ReplyTime                 int       `gorm:"column:reply_time;default:0" json:"reply_time"`
	HmacToken                 string    `gorm:"column:hmac_token" json:"hmac_token"`
	PreChatFormEnabled        bool      `gorm:"column:pre_chat_form_enabled;default:false" json:"pre_chat_form_enabled"`
	PreChatFormOptions        string    `gorm:"column:pre_chat_form_options;type:jsonb;default:'{}'" json:"pre_chat_form_options"`
	HmacMandatory             bool      `gorm:"column:hmac_mandatory;default:false" json:"hmac_mandatory"`
	ContinuityViaEmail        bool      `gorm:"column:continuity_via_email;default:true;not null" json:"continuity_via_email"`
}

// TableName define o nome da tabela para o GORM
func (ChannelWebWidget) TableName() string {
	return "channel_web_widgets"
}

// CreateIndexes cria os índices para a tabela ChannelWebWidget
func (ChannelWebWidget) CreateIndexes(db *gorm.DB) {
	// Criar índice único para o campo hmac_token
	db.Migrator().CreateIndex(&ChannelWebWidget{}, "index_channel_web_widgets_on_hmac_token")

	// Criar índice único para o campo website_token
	db.Migrator().CreateIndex(&ChannelWebWidget{}, "index_channel_web_widgets_on_website_token")
}
