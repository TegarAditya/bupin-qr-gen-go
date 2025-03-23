package utils

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

var font *truetype.Font

func init() {
	var err error
	font, err = loadFont("assets/Poppins-SemiBold.ttf")
	if err != nil {
		panic(err)
	}
}

func clearRect(img *image.RGBA, backgroundColor color.Color, x, y, width, height int) {
	rect := image.Rect(x, y, x+width, y+height)
	draw.Draw(img, rect, &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
}

func loadFont(path string) (*truetype.Font, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return truetype.Parse(fontBytes)
}

func drawText(img *image.RGBA, textStr string, x, y int) {
	const SIZE = 35
	const STROKE_SIZE = 2

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(SIZE)
	c.SetClip(img.Bounds())
	c.SetDst(img)

	c.SetSrc(image.NewUniform(color.White))
	for dx := -STROKE_SIZE; dx <= STROKE_SIZE; dx++ {
		for dy := -STROKE_SIZE; dy <= STROKE_SIZE; dy++ {
			if dx*dx+dy*dy >= STROKE_SIZE*STROKE_SIZE {
				continue
			}
			pt := freetype.Pt(x+dx, y+dy+int(c.PointToFixed(SIZE)>>6-6))
			_, err := c.DrawString(textStr, pt)
			if err != nil {
				panic(err)
			}
		}
	}

	c.SetSrc(image.NewUniform(color.Black))
	pt := freetype.Pt(x, y+int(c.PointToFixed(SIZE)>>6-6))
	_, err := c.DrawString(textStr, pt)
	if err != nil {
		panic(err)
	}
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
