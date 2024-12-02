package migrations

import (
	"log"

	"api/src/models" // ajuste o caminho conforme necessário

	"gorm.io/gorm"
)

// MigrateDatabase realiza a migração dos modelos
func MigrateDatabase(db *gorm.DB) {
	// Migrar os modelos
	err := db.AutoMigrate(
		// Pode ser que tenha que adicionar os campos depois
		&models.Users{},
		&models.Account{},
		&models.AccountUser{},
		&models.AccessToken{},
		&models.ActionMailboxInboundEmail{},
		&models.ActiveStorageAttachment{},
		&models.ActiveStorageBlob{},
		&models.ActiveStorageVariantRecord{},
		&models.AgentBotInbox{},
		&models.AgentBot{},
		&models.AppliedSla{},
		&models.Article{},
		&models.Message{},
		&models.Attachment{},
		&models.Audit{},
		&models.AutomationRule{},
		&models.Campaign{},
		&models.CannedResponse{},
		&models.Category{},
		&models.ChannelAPI{},
		&models.ChannelEmail{},
		&models.ChannelFacebookPage{},
		&models.ChannelLine{},
		&models.ChannelSMS{},
		&models.ChannelTelegram{},
		&models.ChannelTwilioSMS{},
		&models.ChannelTwitterProfile{},
		&models.ChannelWebWidget{},
		&models.ChannelWhatsapp{},
		&models.ContactInbox{},
		&models.Contact{},
		&models.ConversationParticipant{},
		&models.Conversation{},
		&models.CsatSurveyResponse{},
		&models.CustomAttributeDefinition{},
		&models.CustomFilter{},
		&models.CustomRole{},
		&models.DataImport{},
		&models.EmailTemplate{},
		&models.Folder{},
		&models.InboxMember{},
		&models.Inbox{},
		&models.InstallationConfig{},
		&models.IntegrationHook{},
		&models.Label{},
		&models.Macro{},
		&models.Mention{},
		&models.Note{},
		&models.NotificationSetting{},
		&models.NotificationSubscription{},
		&models.Notification{},
		&models.PlatformAppPermissible{},
		&models.PlatformApp{},
		&models.PortalMember{},
		&models.PortalsMember{},
		&models.Portal{},
		&models.RelatedCategory{},
		&models.ReportingEvent{},
		&models.SlaEvent{},
		&models.SlaPolicy{},
		&models.Tagging{},
		&models.Tag{},
		&models.TeamMember{},
		&models.Team{},
		&models.TelegramBot{},
		&models.Webhook{},
		&models.WorkingHour{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar as tabelas: %v", err)
	}

	log.Println("Tabelas migradas com sucesso!")
}