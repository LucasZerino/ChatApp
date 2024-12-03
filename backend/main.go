package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Hello, World!")

	app := fiber.New()

	log.Fatal(app.Listen(":4000"))
}
