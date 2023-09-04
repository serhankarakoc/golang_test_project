package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	SetupRoutes(router fiber.Router)
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	routes := []Router{
		&AuthorRouter{},
	}

	for _, route := range routes {
		route.SetupRoutes(v1)
	}
}
