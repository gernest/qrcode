package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tuotoo/qrcode"
)

func main() {
	fi, err := os.Open("qrcode.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(qrmatrix.Content)
}
