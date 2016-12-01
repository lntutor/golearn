package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	board := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		board[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			if i%2 == 0 {
				board[i][j] = uint8(i * j)
			} else if j%2 == 0 {
				board[i][j] = uint8(i + j)
			} else {
				board[i][j] = uint8((i + j) / 2)
			}
		}
	}
	return board
}

func main() {
	pic.Show(Pic)
}
