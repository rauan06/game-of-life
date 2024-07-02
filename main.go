package main

import (
	"game/assets/flags"
	"game/io"
	"game/logic"
	"time"
)

func main() {
	// Flags handler: flagsMap -> map[string]string
	flagMap := flags.FlagHandler()
	sleepTime := flagMap["delay-ms"].(int)

	// Input handler: grid -> [][]int
	grid := io.MaintainInput(flagMap)

	// Infinite "for" loop
	for {
		io.Show(flagMap, grid)
		logic.CheckState(grid)
		grid = logic.NextGeneration(grid, len(grid), len(grid[0]))
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
}
