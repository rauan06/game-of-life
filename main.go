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
		grid = logic.NextGeneration(grid)
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
}
