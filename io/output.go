package io

import "fmt"

func Help() {
	fmt.Println(`Usage: go run main.go [options]

Options:
  --help        : Show the help message and exit
  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name
  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds
  --file=X      : Load the initial grid from a specified file
  --edges-portal: Enable portal edges where cells that exit the grid appear on the opposite side
  --random=WxH  : Generate a random grid of the specified width (W) and height (H)
  --fullscreen  : Adjust the grid to fit the terminal size with empty cells
  --footprints  : Add traces of visited cells, displayed as '∘'
  --colored     : Add color to live cells and traces if footprints are enabled`)
}

func Show(rules map[string]string, grid [][]int) {

}
