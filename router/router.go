package router

import (
	"bupin-qr-gen-go/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/qrcode/:id", handler.GetQRCode)
}
