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

func main() {
	// Definindo comandos
	migrate := flag.Bool("migrate", false, "Executa migrações")
	seed := flag.Bool("seed", false, "Executa seeders")
	start := flag.Bool("start", false, "Inicia o servidor")

	flag.Parse()

	// Conectar ao banco de dados
	db, err := connectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
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

	if *start {
		log.Println("Iniciando o servidor...")
		startServer(db) // Inicia o servidor
	}
}