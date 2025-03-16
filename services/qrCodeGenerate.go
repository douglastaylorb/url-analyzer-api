package services

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
)

func GenerateQR(url string, includeLogo bool, logoData []byte) (string, error) {
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return "", err
	}

	qrImage := qr.Image(256)

	if includeLogo && len(logoData) > 0 {
		logo, err := imaging.Decode(bytes.NewReader(logoData))
		if err != nil {
			return "", err
		}

		logoSize := qrImage.Bounds().Dx() / 5
		logo = imaging.Resize(logo, logoSize, logoSize, imaging.Lanczos)

		x := (qrImage.Bounds().Dx() - logo.Bounds().Dx()) / 2
		y := (qrImage.Bounds().Dy() - logo.Bounds().Dy()) / 2

		qrImage = imaging.Overlay(qrImage, logo, image.Pt(x, y), 1.0)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, qrImage)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encoded, nil
}
