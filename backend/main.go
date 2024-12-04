package main

import (
	"bufio"
	"log"
	"net"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Inicializa o servidor HTTP com Fiber
	app := fiber.New()

	// Define uma rota GET para o root path "/"
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World 👋!"})
	})

	// Roda o servidor HTTP em uma goroutine
	go func() {
		log.Println("Servidor HTTP rodando na porta 3000")
		log.Fatal(app.Listen(":3000"))
	}()

	// Inicializa o servidor TCP
	startTCPServer(":5000")
}

// Função para iniciar o servidor TCP
func startTCPServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor TCP: %v", err)
	}
	defer listener.Close()

	log.Printf("Servidor TCP rodando na porta %s", port)

	for {
		// Aceita conexões de rastreadores
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Erro ao aceitar conexão: %v", err)
			continue
		}

		log.Printf("Conexão recebida de: %s", conn.RemoteAddr())

		// Processa a conexão em uma goroutine
		go handleTCPConnection(conn)
	}
}

// Função para lidar com conexões TCP
func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		// Lê os dados recebidos do rastreador
		message, err := reader.ReadString('\n') // Se o protocolo não usar '\n', ajuste.
		if err != nil {
			log.Printf("Conexão encerrada por %s: %v", conn.RemoteAddr(), err)
			break
		}

		log.Printf("Dados recebidos de %s: %s", conn.RemoteAddr(), message)

		// Envia um ACK de volta, se necessário
		ack := "ACK\n"
		conn.Write([]byte(ack))
	}
}