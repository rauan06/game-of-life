package assets

func CountLiveCells(grid [][]int) int {
	counter := 0
	for _, r := range grid {
		for _, s := range r {
			if s == 1 {
				counter++
			}
		}
	}

	return counter
}
