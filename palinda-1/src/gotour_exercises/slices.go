package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	grid := make([][]uint8, dy)
	for i := range grid {
		grid[i] = make([]uint8, dx)
		for j := range dx {
			grid[i][j] = uint8((j + i) / 2)
		}

	}
	return grid
}
func slices() {
	println("slices exercise:")
	pic.Show(Pic)
}
