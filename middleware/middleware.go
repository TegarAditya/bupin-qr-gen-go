package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type QRCodeRequest struct {
	ID string `validate:"required,startswith=UJN-|startswith=VID-"`
}

func ValidateQRCodeRequest(c *fiber.Ctx) error {
	id := c.Params("id")

	req := QRCodeRequest{ID: id}
	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Next()
}
