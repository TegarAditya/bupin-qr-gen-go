package utils

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"sync"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func clearRect(img *image.RGBA, backgroundColor color.Color, x, y, width, height int) {
	rect := image.Rect(x, y, x+width, y+height)
	draw.Draw(img, rect, &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
}

func drawText(img *image.RGBA, text string, x, y int) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{
		X: fixed.Int26_6(x << 6),
		Y: fixed.Int26_6(y << 6),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}

func GenerateQRCode(data string) ([]byte, error) {
	const MARK_WIDTH = 145
	const MARK_HEIGHT = 39
	const QR_SIZE = 512

	qrcode, err := qr.Encode(data, qr.Q, qr.Auto)
	if err != nil {
		return nil, err
	}

	qrcode, err = barcode.Scale(qrcode, QR_SIZE, QR_SIZE)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(qrcode.Bounds())
	draw.Draw(img, img.Bounds(), qrcode, image.Point{}, draw.Src)

	clearRect(img, color.Transparent, 351, 457, MARK_WIDTH, MARK_HEIGHT)

	drawText(img, "bupin.id", 351, 457)

	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err = png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
