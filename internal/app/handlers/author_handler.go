package handlers

import (
	"strconv"

	"starter/helpers"
	"starter/internal/app/dtos"
	"starter/internal/app/repositories"
	"starter/internal/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func validateAuthorInputs(input interface{}) error {
	validate := validator.New()

	return validate.Struct(input)
}

func authorService() *services.AuthorService {
	repo := repositories.NewAuthorRepository()

	return services.NewAuthorService(repo)
}

func GetAllAuthors(c *fiber.Ctx) error {
	service := authorService()

	authors, err := service.GetAllAuthors()

	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", authors)
}

func GetAuthorByID(c *fiber.Ctx) error {
	authorIDParam := c.Params("id")

	authorID, err := strconv.Atoi(authorIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid author ID", nil))
	}

	service := authorService()

	author, err := service.GetAuthorByID(uint(authorID))

	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", author)
}

func CreateAuthor(c *fiber.Ctx) error {
	var input dtos.CreateAuthorDTO

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateAuthorInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	service := authorService()

	createdAuthor, err := service.CreateAuthor(input)

	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", createdAuthor)
}

func UpdateAuthor(c *fiber.Ctx) error {
	authorIDParam := c.Params("id")

	authorID, err := strconv.Atoi(authorIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid author ID", nil))
	}

	var input dtos.UpdateAuthorDTO

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid request payload", nil))
	}

	if err := validateAuthorInputs(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid input data", nil))
	}

	service := authorService()

	updatedAuthor, err := service.UpdateAuthor(uint(authorID), input)

	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.SendJSONResponse(c, fiber.StatusOK, "Success", updatedAuthor)
}

func DeleteAuthor(c *fiber.Ctx) error {
	authorIDParam := c.Params("id")

	authorID, err := strconv.Atoi(authorIDParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.NewResponse(fiber.StatusBadRequest, "Invalid author ID", nil))
	}

	service := authorService()

	delErr := service.DeleteAuthor(uint(authorID))

	if delErr != nil {
		return helpers.HandleError(c, delErr)
	}

	return helpers.SendJSONResponse(c, fiber.StatusNoContent, "Success", nil)
}
