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

//TODO Configurable
const (
	COLOR_FG_START           = 0 //foreground
	COLOR_FG_END             = 150
	COLOR_BG_START           = 200 //background
	COLOR_BG_END             = 255
	FONT_SIZE_DRIFT_DECREASE = 3 // actual_size >= base_size - FONT_SIZE_DRIFT_DECREASE
	FONT_SIZE_DRIFT_INCREASE = 5 // actual_size <= base_size + FONT_SIZE_DRIFT_INCREASE
)

var random = rand.New(rand.NewSource(time.Now().Unix() + 123569))

func CreatePng(fontFile, data string, size, dpi, width, height int) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.NewUniform(randomBGColor()), image.ZP, draw.Src)
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	faceOpt := &truetype.Options{
		Size:    float64(size),
		DPI:     float64(dpi),
		Hinting: font.HintingFull,
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: randomFGColor()},
		Face: truetype.NewFace(f, faceOpt),
	}
	y := (height + size) / 2
	d.Dot = fixed.Point26_6{
		X: (fixed.I(width) - d.MeasureString(data)) / 2,
		Y: fixed.I(y),
	}
	for _, c := range data {
		d.DrawString(string(c))
		d.Src = image.NewUniform((randomFGColor()))
		faceOpt.Size = randomFontSize(size)
		d.Face = truetype.NewFace(f, faceOpt)
	}
	if err != nil {
		return nil, err
	}
	return img, nil

}

//s:start e:end
func randomNum(s, e int) int {
	return s + random.Intn(e-s)
}

func randomColorNum(s, e int) uint8 {
	return uint8(randomNum(s, e))
}

func randomFGColor() color.Color {
	c := &color.RGBA{
		R: randomColorNum(COLOR_FG_START, COLOR_FG_END),
		G: randomColorNum(COLOR_FG_START, COLOR_FG_END),
		B: randomColorNum(COLOR_FG_START, COLOR_FG_END),
		A: 255,
	}
	return c
}

func randomBGColor() color.Color {
	c := &color.RGBA{
		R: randomColorNum(COLOR_BG_START, COLOR_BG_END),
		G: randomColorNum(COLOR_BG_START, COLOR_BG_END),
		B: randomColorNum(COLOR_BG_START, COLOR_BG_END),
		A: 255,
	}
	return c
}

func randomFontSize(baseSize int) float64 {
	return float64(baseSize - FONT_SIZE_DRIFT_DECREASE + randomNum(0, FONT_SIZE_DRIFT_DECREASE+FONT_SIZE_DRIFT_INCREASE))
}
