package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	color         uint8
}

func (self *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.Width, self.Height)
}

func (self *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self *Image) At(x, y int) color.Color {
	return color.RGBA{self.color + uint8(x), self.color + uint8(y), 255, 255}
}

func main() {
	image := Image{128, 128, 128}
	pic.ShowImage(&image)
}