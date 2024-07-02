package logic

func NextGeneration(grid [][]int, h int, w int) [][]int {
	future := make([][]int, h)
	for i := range future {
		future[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			aliveNeighbours := 0

			// Iterate over neighbors
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					// Skip the cell itself
					if di == 0 && dj == 0 {
						continue
					}
					// Check boundaries
					if (i+di >= 0 && i+di < h) && (j+dj >= 0 && j+dj < w) {
						aliveNeighbours += grid[i+di][j+dj]
					}
				}
			}

			// Apply rules of life
			if grid[i][j] == 1 && (aliveNeighbours < 2 || aliveNeighbours > 3) {
				future[i][j] = 0 // Cell dies due to underpopulation or overpopulation
			} else if grid[i][j] == 0 && aliveNeighbours == 3 {
				future[i][j] = 1 // A new cell is born
			} else {
				future[i][j] = grid[i][j] // Remains the same
			}
		}
	}

	return future
}
