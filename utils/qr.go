package utils

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	MARK_WIDTH  = 145
	MARK_HEIGHT = 39
	QR_SIZE     = 512
)

var (
	font *truetype.Font
	ctx  *freetype.Context
)

func init() {
	var err error
	font, err = loadFont("assets/Poppins-SemiBold.ttf")
	if err != nil {
		panic(err)
	}

	ctx = freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(font)
	ctx.SetFontSize(35)
}

func loadFont(path string) (*truetype.Font, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return truetype.Parse(fontBytes)
}

func drawText(img *image.RGBA, textStr string, x, y int) {
	const (
		SIZE        = 35
		STROKE_SIZE = 2
	)

	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)

	fixedSize := int(ctx.PointToFixed(SIZE) >> 6)
	baseY := y + fixedSize - 6

	// Draw stroke
	ctx.SetSrc(image.NewUniform(color.White))
	for dx := -STROKE_SIZE; dx <= STROKE_SIZE; dx++ {
		for dy := -STROKE_SIZE; dy <= STROKE_SIZE; dy++ {
			if dx*dx+dy*dy >= STROKE_SIZE*STROKE_SIZE {
				continue
			}
			pt := freetype.Pt(x+dx, baseY+dy)
			_, _ = ctx.DrawString(textStr, pt)
		}
	}

	// Draw text
	ctx.SetSrc(image.NewUniform(color.Black))
	_, _ = ctx.DrawString(textStr, freetype.Pt(x, baseY))
}

func drawWatermark(img *image.RGBA) {
	markRect := image.Rect(351, 457, 351+MARK_WIDTH, 457+MARK_HEIGHT)
	draw.Draw(img, markRect, image.NewUniform(color.Transparent), image.Point{}, draw.Src)

	drawText(img, "bupin.id", 351, 457)
}

func GenerateQRCode(data string, watermark bool) ([]byte, error) {
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

	if watermark {
		drawWatermark(img)
	}

	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
