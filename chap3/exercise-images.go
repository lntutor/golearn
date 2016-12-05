package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math"
)

type Image struct {
	width  int
	height int
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img *Image) At(x, y int) color.Color {
	tmp1 := uint8(x + y)
	tmp2 := uint8(x * y)
	tmp3 := uint8((x + y) / 2)
	tmp4 := uint8(math.Abs(float64(x-y)) / 2)
	return color.RGBA{tmp1, tmp2, tmp3, tmp4}
}

func main() {
	m := Image{222, 333}
	pic.ShowImage(&m)
}
