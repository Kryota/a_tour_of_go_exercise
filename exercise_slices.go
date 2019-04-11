package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i, _ := range y {
		y[i] = make([]uint8, dx)
		for j, _ := range y[i] {
			y[i][j] = uint8(dx * dy)
		}
	}
	return y
}

func main() {
	pic.Show(Pic)
}
