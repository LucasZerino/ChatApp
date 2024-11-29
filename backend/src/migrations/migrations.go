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
		&models.Users{},
		&models.Account{},
		&models.AccountUser{},
		&models.TeamMember{},
		&models.PortalMember{},
		&models.DataImport{},
		&models.CustomRole{},
		&models.EmailTemplate{},
		&models.Contact{},
		&models.Tag{},
		&models.AgentBot{},
		&models.ConversationParticipant{},
		&models.ChannelWebWidget{},
		&models.Inbox{},
		&models.Message{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar as tabelas: %v", err)
	}

	log.Println("Tabelas migradas com sucesso!")
}