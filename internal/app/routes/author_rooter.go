package routes

import (
	"starter/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type AuthorRouter struct{}

func (authorRouter *AuthorRouter) SetupRoutes(router fiber.Router) {
	authors := router.Group("/authors")
	authors.Get("/", handlers.GetAllAuthors)
	authors.Get("/:id", handlers.GetAuthorByID)
	authors.Post("/", handlers.CreateAuthor)
	authors.Put("/:id", handlers.UpdateAuthor)
	authors.Delete("/:id", handlers.DeleteAuthor)
}
