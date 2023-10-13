package server

import (
	"log"

	"starter/internal/app/routes"
	"starter/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Server() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitLogger()
	defer utils.GetLogger().Sync()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
