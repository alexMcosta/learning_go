package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Test")
	qrcode := GenerateQRCode("555-2368")

	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

//GenerateQRCode will be the function that generates the QR Code
func GenerateQRCode(code string) []byte {
	return []byte{0xFF}
}
