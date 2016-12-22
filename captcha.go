package gocaptcha

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	COLOR_FG_START = 0 //foreground
	COLOR_FG_END   = 150
	COLOR_BG_START = 200 //background
	COLOR_BG_END   = 255
)

var random = rand.New(rand.NewSource(time.Now().Unix() + 123569))

func CreatePng(fontFile, data string, size, dpi, width, height int) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), getUniform(RandomBGColor()), image.ZP, draw.Src)
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
		Src: &image.Uniform{C: RandomFGColor()},
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
	for _, c := range data {
		d.DrawString(string(c))
		d.Src = getUniform(RandomFGColor())
	}
	if err != nil {
		return nil, err
	}
	return img, nil

}

func getUniform(c color.Color) *image.Uniform {
	return &image.Uniform{C: c}
}

func RandomFGColor() color.Color {
	c := &color.RGBA{
		R: randomColorNum(COLOR_FG_START, COLOR_FG_END),
		G: randomColorNum(COLOR_FG_START, COLOR_FG_END),
		B: randomColorNum(COLOR_FG_START, COLOR_FG_END),
		A: 255,
	}
	return c
}

func RandomBGColor() color.Color {
	c := &color.RGBA{
		R: randomColorNum(COLOR_BG_START, COLOR_BG_END),
		G: randomColorNum(COLOR_BG_START, COLOR_BG_END),
		B: randomColorNum(COLOR_BG_START, COLOR_BG_END),
		A: 255,
	}
	return c
}

//s:start e:end
func randomColorNum(s, e int) uint8 {
	return uint8(s + random.Intn(e-s))
}
