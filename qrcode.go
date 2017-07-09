package main

import (
	"fmt"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func genqrcodepng(filename string, info string, level qrcode.RecoveryLevel) bool {

	err := qrcode.WriteFile(info, qrcode.Medium, 256, filename)

	if err != nil {
		err = os.Remove(filename)
	}

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	return true
}
