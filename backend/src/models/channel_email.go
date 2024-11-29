package models

import (
	"time"

	"gorm.io/gorm"
)

// ChannelEmail representa a tabela channel_email
type ChannelEmail struct {
	ID                     uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID              uint                   `gorm:"column:account_id;not null;index" json:"account_id"`
	Email                  string                 `gorm:"column:email;not null;uniqueIndex" json:"email"`
	ForwardToEmail         string                 `gorm:"column:forward_to_email;not null;uniqueIndex" json:"forward_to_email"`
	CreatedAt              time.Time              `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt              time.Time              `gorm:"column:updated_at;not null" json:"updated_at"`
	IMAPEnabled            bool                   `gorm:"column:imap_enabled;default:false" json:"imap_enabled"`
	IMAPAddress            string                 `gorm:"column:imap_address;default:''" json:"imap_address"`
	IMAPPort               int                    `gorm:"column:imap_port;default:0" json:"imap_port"`
	IMAPLogin              string                 `gorm:"column:imap_login;default:''" json:"imap_login"`
	IMAPPassword           string                 `gorm:"column:imap_password;default:''" json:"imap_password"`
	IMAPEnableSSL          bool                   `gorm:"column:imap_enable_ssl;default:true" json:"imap_enable_ssl"`
	SMTPEnabled            bool                   `gorm:"column:smtp_enabled;default:false" json:"smtp_enabled"`
	SMTPAddress            string                 `gorm:"column:smtp_address;default:''" json:"smtp_address"`
	SMTPPort               int                    `gorm:"column:smtp_port;default:0" json:"smtp_port"`
	SMTPLogin              string                 `gorm:"column:smtp_login;default:''" json:"smtp_login"`
	SMTPPassword           string                 `gorm:"column:smtp_password;default:''" json:"smtp_password"`
	SMTPDomain             string                 `gorm:"column:smtp_domain;default:''" json:"smtp_domain"`
	SMTPEnableStartTLSAuto bool                   `gorm:"column:smtp_enable_starttls_auto;default:true" json:"smtp_enable_starttls_auto"`
	SMTPAuthentication     string                 `gorm:"column:smtp_authentication;default:'login'" json:"smtp_authentication"`
	SMTPOpenSSLVerifyMode  string                 `gorm:"column:smtp_openssl_verify_mode;default:'none'" json:"smtp_openssl_verify_mode"`
	SMTPEnableSSLTLS       bool                   `gorm:"column:smtp_enable_ssl_tls;default:false" json:"smtp_enable_ssl_tls"`
	ProviderConfig         map[string]interface{} `gorm:"column:provider_config;default:'{}'" json:"provider_config"`
	Provider              string                 `gorm:"column:provider" json:"provider"`
}

// TableName define o nome da tabela para o GORM
func (ChannelEmail) TableName() string {
	return "channel_email"
}

// CreateIndexes cria os índices para a tabela ChannelEmail
func (ChannelEmail) CreateIndexes(db *gorm.DB) {
	// Criar índices únicos
	db.Migrator().CreateIndex(&ChannelEmail{}, "index_channel_email_on_email")
	db.Migrator().CreateIndex(&ChannelEmail{}, "index_channel_email_on_forward_to_email")
}
