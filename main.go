package main

import (
	"game/assets/flags"
	"game/io"
	"game/logic"
	"time"
)

func main() {
	// Flags handler: flagsMap -> map[string]interface{}
	flagMap := flags.FlagHandler()
	sleepTime := flagMap["delay-ms"].(int)

	// Input handler: grid -> [][]int
	grid := io.MaintainInput(flagMap)

	// Infinite "for" loop
	tick := 0
	for {
		tick++
		io.Show(flagMap, grid, tick)
		logic.CheckState(grid)
		grid = logic.NextGeneration(
			grid,
			len(grid),
			len(grid[0]),
			flagMap["edges-portal"].(bool),
		)
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
}
