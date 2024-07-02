package assets

import "math/rand"

func RandomGrid(h, w int) [][]int {
	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			grid[i][j] = rand.Intn(2)
		}
	}

	return grid
}
