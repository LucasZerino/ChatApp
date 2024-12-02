package main

import (
	"flag"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/src/controllers" // ajuste o caminho conforme necessário
	"api/src/migrations"  // ajuste o caminho conforme necessário
	"api/src/seeders"     // adicione esta linha para importar o pacote de seeders

	"github.com/gin-gonic/gin"
)

// Função para conectar ao banco de dados
func connectDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable" // ajuste conforme necessário
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Função para semear dados
func seedDatabase(db *gorm.DB) {
	
}

func startServer(db *gorm.DB) {
	r := gin.Default()
	r.POST("/login/superadmin", controllers.LoginSuperAdmin(db)) // Passa a instância do db

	if err := r.Run(":8080"); err != nil { // escuta na porta 8080
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func deleteDatabase(db *gorm.DB) {
	// Aqui você pode usar um comando SQL para deletar o banco de dados
	if err := db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error; err != nil {
		log.Fatalf("Erro ao deletar o banco de dados: %v", err)
	}
	log.Println("Banco de dados deletado com sucesso!")
}

// Função para limpar o banco de dados
func clearDatabase(db *gorm.DB) {
	// Obter todas as tabelas
	tables := []string{
		"users",
		"accounts",
		"access_tokens",
		"action_mailbox_inbound_emails",
		"active_storage_attachments",
		"active_storage_blobs",
		"active_storage_variant_records",
		"agent_bots_inbox",
		"agent_bots",
		"applied_slas",
		"articles",
		"attachments",
		"audits",
		"automation_rules",
		"campaigns",
		"canned_responses",
		"categories",
		"channel_api",
		"channel_email",
		"channel_facebook_pages",
		"channel_line",
		"channel_sms",
		"channel_telegram",
		"channel_twilio_sms",
		"channel_twitter_profiles",
		"channel_web_widgets",
		"channel_whatsapp",
		"contact_inbox",
		"contacts",
		"conversation_participants",
		"conversations",
		"csat_survey_responses",
		"custom_attribute_definitions",
		"custom_filters",
		"custom_roles",
		"data_imports",
		"email_templates",
		"folders",
		"inbox_members",
		"inboxes",
		"installation_configs",
		"integration_hooks",
		"labels",
		"macros",
		"mentions",
		"messages",
		"notes",
		"notification_settings",
		"notification_subscriptions",
		"notifications",
		"platform_app_permissibles",
		"platform_apps",
		"portal_members",
		"portals_members",
		"portals",
		"related_categories",
		"reporting_events",
		"sla_events",
		"sla_policies",
		"taggings",
		"tags",
		"team_members",
		"teams",
		"telegram_bots",
		"webhooks",
		"working_hours",
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Printf("Erro ao limpar a tabela %s: %v", table, err)
		} else {
			log.Printf("Tabela %s removida com sucesso.", table)
		}
	}
}

func main() {
	// Definindo comandos
	migrate := flag.Bool("migrate", false, "Executa migrações")
	seed := flag.Bool("seed", false, "Executa seeders")
	clear := flag.Bool("clear", false, "Limpa o banco de dados")
	start := flag.Bool("start", false, "Inicia o servidor")
	delete := flag.Bool("delete", false, "Deleta o banco de dados") // Adicionando o flag para deletar o banco de dados

	flag.Parse()

	// Conectar ao banco de dados
	db, err := connectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if *delete {
		deleteDatabase(db) // Chama a função para deletar o banco de dados
		return
	}

	if *migrate {
		log.Println("Executando migrações...")
		migrations.MigrateDatabase(db) // Chama a função de migração
	}

	if *seed {
		log.Println("Executando seed...")
		seeders.SeedDatabase(db) // Chama a função SeedDatabase do pacote seeders
		log.Println("Dados semeados com sucesso!")
	}

	if *clear {
		log.Println("Limpando o banco de dados...")
		clearDatabase(db) // Chama a função de limpeza
	}

	if *start {
		log.Println("Iniciando o servidor...")
		startServer(db) // Inicia o servidor
	}
}