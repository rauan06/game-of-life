package main

import (
	"game/assets/flags"
	"game/io"
	"game/logic"
)

func main() {
	// Flags handler: flagsMap -> map[string]string
	flagMap := flags.FlagHandler()

	// Input handler: grid -> [][]int
	grid := io.MaintainInput(flagMap)

	// Infinite "for" loop
	for {
		io.Show(flagMap, grid)
		grid = logic.NextGeneration(grid)
	}
}
