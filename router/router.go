package router

import (
	"bupin-qr-gen-go/handler"
	"bupin-qr-gen-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/qrcode/:id", middleware.ValidateQRCodeRequest, handler.GetQRCode)
}
