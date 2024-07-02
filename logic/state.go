package logic

import (
	"fmt"
	"os"
)

func CheckState(grid [][]int) {
	for _, r := range grid {
		for _, t := range r {
			if t == 1 {
				return
			}
		}
	}

	fmt.Println("No live cell left")
	os.Exit(0)
}
