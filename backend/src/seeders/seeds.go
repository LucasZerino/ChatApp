package seeders

import (
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/src/models" // ajuste o caminho conforme necessário
)


func ptr(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}

// Função para conectar ao banco de dados
func connectDatabase() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable" // ajuste conforme necessário
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Função para semear dados
func SeedDatabase(db *gorm.DB) {
	log.Println("Iniciando o seeding do banco de dados...")

	// Criação de contas
	account := models.Account{Name: "Acme Inc"}
	if err := db.Create(&account).Error; err != nil {
		log.Fatalf("Erro ao criar conta: %v", err)
	}
	log.Println("Conta criada:", account)

	secondaryAccount := models.Account{Name: "Acme Org"}
	if err := db.Create(&secondaryAccount).Error; err != nil {
		log.Fatalf("Erro ao criar conta secundária: %v", err)
	}
	log.Println("Conta secundária criada:", secondaryAccount)

	// Criação de usuário
	user := models.Users{
		Name:              "Beto",
		Email:             ptr("beto@bolt360.com.br"),
		EncryptedPassword: "Password1!",
		Type:              ptr("SuperAdmin"),
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Erro ao criar usuário: %v", err)
	}
	log.Println("Usuário criado:", user)

	// Criação de associações de conta e usuário
	if err := db.Create(&models.AccountUser{
		AccountID: account.ID,
		UserID:    int64(user.ID),
		Role:      1,
	}).Error; err != nil {
		log.Fatalf("Erro ao criar AccountUser: %v", err)
	}

	if err := db.Create(&models.AccountUser{
		AccountID: secondaryAccount.ID,
		UserID:    int64(user.ID),
		Role:      1,
	}).Error; err != nil {
		log.Fatalf("Erro ao criar AccountUser secundário: %v", err)
	}

	// Criação de outros dados
	webWidget := models.ChannelWebWidget{AccountID: uint(account.ID), WebsiteURL: "https://acme.inc"}
	if err := db.Create(&webWidget).Error; err != nil {
		log.Fatalf("Erro ao criar WebWidget: %v", err)
	}

	inbox := models.Inbox{ChannelID: webWidget.ID, AccountID: uint(account.ID), Name: "Acme Support"}
	if err := db.Create(&inbox).Error; err != nil {
		log.Fatalf("Erro ao criar Inbox: %v", err)
	}

	// Adicionando membro à Inbox
	if err := db.Create(&models.InboxMember{
		UserID: uint(user.ID),
		InboxID: inbox.ID,
	}).Error; err != nil {
		log.Fatalf("Erro ao criar InboxMember: %v", err)
	}

	// Criação de conversa
	contactInbox := models.ContactInbox{
		SourceID:        strconv.Itoa(int(user.ID)),
		InboxID:        inbox.ID,
		HmacVerified:    true,
		Contact: models.Contact{Name: "jane", Email: "jane@example.com", PhoneNumber: "+2320000"},
	}
	if err := db.Create(&contactInbox).Error; err != nil {
		log.Fatalf("Erro ao criar ContactInbox: %v", err)
	}

	conversation := models.Conversation{
		AccountID: uint(account.ID),
		InboxID:   inbox.ID,
		Status:    1,
		AssigneeID: &user.ID,
		ContactID: &contactInbox.ContactID,
	}
	if err := db.Create(&conversation).Error; err != nil {
		log.Fatalf("Erro ao criar Conversation: %v", err)
	}

	// Mensagens de exemplo
	// Mensagem de email
	if err := db.Create(&models.Message{
		Content:    "Hello",
		AccountID: uint(account.ID),
		InboxID:   inbox.ID,
		ConversationID: conversation.ID,
		SenderID:  &contactInbox.ContactID,
		MessageType: 1,
	}).Error; err != nil {
		log.Fatalf("Erro ao criar mensagem: %v", err)
	}

	// Mensagem de localização
	locationMessage := models.Message{
		Content:    "location",
		AccountID: uint(account.ID),
		InboxID:   inbox.ID,
		SenderID:  &contactInbox.ContactID,
		ConversationID: conversation.ID,
		MessageType:    1,
	}
	locationMessage.Attachments = append(locationMessage.Attachments, models.Attachment{
		AccountID: int(account.ID),
		FileType:  1,
		CoordinatesLat: 37.7893768,
		CoordinatesLong: -122.3895553,
		FallbackTitle: ptr("Bay Bridge, San Francisco, CA, USA"),
	})
	if err := db.Create(&locationMessage).Error; err != nil {
		log.Fatalf("Erro ao criar mensagem de localização: %v", err)
	}

	// Respostas automáticas
	if err := db.Create(&models.CannedResponse{
		AccountID: uint(account.ID),
		ShortCode: "start",
		Content:   "Hello welcome to bchat.",
	}).Error; err != nil {
		log.Fatalf("Erro ao criar CannedResponse: %v", err)
	}
}

func main() {
	// Conectar ao banco de dados
	db, err := connectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Semear dados
	SeedDatabase(db)

	log.Println("Dados semeados com sucesso!")
}