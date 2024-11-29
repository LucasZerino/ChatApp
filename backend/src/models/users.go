package models

import (
	"time"

	"gorm.io/datatypes"
)

// User representa a tabela users
type Users struct {
	ID                     uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Provider               string         `gorm:"column:provider;default:'email';not null" json:"provider"`
	UID                    string         `gorm:"column:uid;default:'';not null" json:"uid"`
	EncryptedPassword      string         `gorm:"column:encrypted_password;default:'';not null" json:"encrypted_password"`
	ResetPasswordToken     *string        `gorm:"column:reset_password_token" json:"reset_password_token"`
	ResetPasswordSentAt    *time.Time     `gorm:"column:reset_password_sent_at" json:"reset_password_sent_at"`
	RememberCreatedAt      *time.Time     `gorm:"column:remember_created_at" json:"remember_created_at"`
	SignInCount            int            `gorm:"column:sign_in_count;default:0;not null" json:"sign_in_count"`
	CurrentSignInAt        *time.Time     `gorm:"column:current_sign_in_at" json:"current_sign_in_at"`
	LastSignInAt           *time.Time     `gorm:"column:last_sign_in_at" json:"last_sign_in_at"`
	CurrentSignInIP        *string        `gorm:"column:current_sign_in_ip" json:"current_sign_in_ip"`
	LastSignInIP           *string        `gorm:"column:last_sign_in_ip" json:"last_sign_in_ip"`
	ConfirmationToken      *string        `gorm:"column:confirmation_token" json:"confirmation_token"`
	ConfirmedAt            *time.Time     `gorm:"column:confirmed_at" json:"confirmed_at"`
	ConfirmationSentAt     *time.Time     `gorm:"column:confirmation_sent_at" json:"confirmation_sent_at"`
	UnconfirmedEmail       *string        `gorm:"column:unconfirmed_email" json:"unconfirmed_email"`
	Name                   string         `gorm:"column:name;not null" json:"name"`
	DisplayName            *string        `gorm:"column:display_name" json:"display_name"`
	Email                  *string        `gorm:"column:email" json:"email"`
	Tokens                 datatypes.JSON `gorm:"column:tokens" json:"tokens"`
	CreatedAt              time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	PubsubToken            *string        `gorm:"column:pubsub_token" json:"pubsub_token"`
	Availability           int            `gorm:"column:availability;default:0" json:"availability"`
	UISettings             datatypes.JSON `gorm:"column:ui_settings;default:'{}'" json:"ui_settings"`
	CustomAttributes       datatypes.JSON `gorm:"column:custom_attributes;default:'{}'" json:"custom_attributes"`
	Type                   *string        `gorm:"column:type" json:"type"`
	MessageSignature       *string        `gorm:"column:message_signature" json:"message_signature"`
}

// TableName define o nome da tabela para o GORM
func (Users) TableName() string {
	return "users"
}
