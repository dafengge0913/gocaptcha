package gocaptcha

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"io/ioutil"
)

func CreatePng(fontFile, data string, size, dpi, width, height int) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bgColor := color.RGBA{
		R: 250,
		G: 250,
		B: 240,
		A: 255,
	}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	d := &font.Drawer{
		Dst: img,
		Src: image.Black,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    float64(size),
			DPI:     float64(dpi),
			Hinting: font.HintingFull,
		}),
	}
	y := (height + size) / 2
	d.Dot = fixed.Point26_6{
		X: (fixed.I(width) - d.MeasureString(data)) / 2,
		Y: fixed.I(y),
	}
	d.DrawString(data)
	if err != nil {
		return nil, err
	}
	return img, nil

}
