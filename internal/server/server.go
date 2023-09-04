package server

import (
	"log"

	"zatrano/internal/app/routes"
	"zatrano/internal/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Server() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	helpers.InitLogger()
	defer helpers.GetLogger().Sync()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
