package main

import (
	"fmt"
	"github.com/dafengge0913/gocaptcha"
	"image/png"
	"os"
)

func main() {
	img, err := gocaptcha.CreatePng("example/arial.ttf", "9527", 40, 72, 100, 50)
	if err != nil {
		fmt.Println("create error :", err)
		return
	}

	file, err := os.OpenFile("example/png/output/1.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("OpenFile error :", err)
		return
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Encode error :", err)
		return
	}

}
