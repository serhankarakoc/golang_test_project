package helpers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func HandleError(c *fiber.Ctx, err error) error {
	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(NewResponse(fiber.StatusNotFound, "Not found", nil))
	}
	return c.Status(fiber.StatusInternalServerError).JSON(NewResponse(fiber.StatusInternalServerError, "Server error", nil))
}

func SendJSONResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.JSON(NewResponse(code, message, data))
}
