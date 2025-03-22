package utils

import (
	"bytes"
	"image/png"
	"sync"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GenerateQRCode(data string) ([]byte, error) {
	qrCode, err := qr.Encode(data, qr.Q, qr.Auto)
	if err != nil {
		return nil, err
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		return nil, err
	}

	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err = png.Encode(buf, qrCode)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
