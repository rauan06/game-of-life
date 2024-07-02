package io

import (
	"bufio"
	"fmt"
	"game/assets"
	"os"
	"strconv"
	"strings"
)

func MaintainInput(flagsMap map[string]interface{}) [][]int {
	if flagsMap["file"] != false {
		return assets.Reader(flagsMap["file"].(string))
	}

	if flagsMap["random"] == false {
		return InputHandler()
	}
	return assets.RandomGrid(flagsMap["random"].([]int)[1], flagsMap["random"].([]int)[0])
}

func InputHandler() [][]int {
	var h, w int

	reader := bufio.NewReader(os.Stdin)

	for {
		var err error
		fmt.Print("Enter the size of the grid (height x width): ")

		firstLine, _ := reader.ReadString('\n')
		firstLine = strings.TrimSpace(firstLine)

		firstLineArgs := strings.SplitN(firstLine, " ", 2)

		h, err = strconv.Atoi(firstLineArgs[0])

		if err != nil {
			fmt.Println("invalid input, please try again")
			continue
		}

		w, err = strconv.Atoi(firstLineArgs[1])

		if err != nil {
			fmt.Println("invalid input, please try again")
			continue
		}

		break
	}

	grid := make([][]int, h)
	for r := range grid {
		grid[r] = make([]int, w)
	}

	return readGrid(h, w)
}

func readGrid(h, w int) [][]int {
	reader := bufio.NewReader(os.Stdin)
	grid := make([][]int, h)

	for r := 0; r < h; r++ {
		grid[r] = make([]int, w)
	}

	counter := 1
	for {
		if counter > h {
			break
		}
		fmt.Printf("Enter row %d: ", counter)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if len([]rune(line)) == w {
			arr := []int{}

			for _, r := range line {
				if r == '.' {
					arr = append(arr, 0)
				} else {
					arr = append(arr, 1)
				}
			}

			grid[counter-1] = arr
			counter++
		} else {
			fmt.Println("Invalid line length, please try again")
			continue
		}
	}

	return grid
}
