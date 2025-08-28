package handler

import (
	"bupin-qr-gen-go/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetQRCode(c *fiber.Ctx) error {
	id := c.Params("id")

	var (
		value any
		err   error
	)

	if strings.HasPrefix(id, "UJN") {
		value, err = utils.GetInfoUJN(id)
	} else {
		value, err = utils.GetInfoVID(id)
	}

	if err != nil {
		return c.Status(404).SendString("String not found")
	}

	filename := utils.GetFileName(value)
	url := "https://buku.bupin.id/?" + id

	qrCode, err := utils.GenerateQRCode(url, true)
	if err != nil {
		return c.Status(500).SendString("Failed to generate QR code")
	}

	c.Set("Content-Type", "image/png")
	c.Set("Content-Disposition", "inline; filename="+filename+".png")
	return c.Send(qrCode)
}
