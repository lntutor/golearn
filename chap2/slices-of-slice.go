package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

}

func main() {
	pic.Show(Pic)
}
