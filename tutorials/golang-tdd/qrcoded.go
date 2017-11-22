package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
)

func main() {
	fmt.Println("Test")
	qrcode := GenerateQRCode("555-2368")

	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

//GenerateQRCode will be the function that generates the QR Code
func GenerateQRCode(code string) []byte {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)

	return buf.Bytes()
}
