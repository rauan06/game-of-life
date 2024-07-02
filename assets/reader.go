package assets

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Reader(filename string) [][]int {
	var h, w int

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	// Parse height and width from the first line
	if len(lines) < 1 {
		log.Fatal("Invalid file format: missing height and width")
	}
	firstLine := strings.TrimSpace(lines[0])
	parts := strings.Split(firstLine, " ")
	if len(parts) != 2 {
		log.Fatal("Invalid file format: expected height and width")
	}

	h, err = strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Error converting height value: %v", err)
	}

	w, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error converting width value: %v", err)
	}

	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}

	// Parse grid data from subsequent lines
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		for j, r := range line {
			if string(r) == "#" {
				grid[i-1][j] = 1
			} else if string(r) == "." {
				grid[i-1][j] = 0
			} else {
				log.Fatal("invalid symbol in a file")
			}
		}
	}

	return grid
}
