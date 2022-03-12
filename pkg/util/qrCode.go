package util

import (
	qrcode "github.com/skip2/go-qrcode"
	"log"
)

func generateQRCode(url string) {
	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qr.png")
	if err != nil {
		log.Fatalf("QRCode generate fail: %v", err)
	}
}
