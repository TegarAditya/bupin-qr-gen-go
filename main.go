package main

import (
	"bupin-qr-gen-go/config"
	"bupin-qr-gen-go/database"
	"bupin-qr-gen-go/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.LoadConfig()

	database.InitDB()

	app := fiber.New()

	router.SetupRoutes(app)

	log.Printf("Starting server on port %d", cfg.Port)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
