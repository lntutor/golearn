package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 10, 10))
	fmt.Println(m.Bounds())
	fmt.Println(m.ColorModel())
}
