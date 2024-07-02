package io

import (
	"fmt"
	"game/assets"
	"log"
)

func MaintainInput(flagsMap map[string]interface{}) [][]int {
	if flagsMap["random"] == false {
		temp, err := InputHandler()

		if err != nil {
			log.Fatal(err)
		}

		return temp
	}
	return assets.RandomGrid()
}

func InputHandler() ([][]int, error) {
	var w, h int
	fmt.Print("Enter the size of the grid (width x height): ")
	fmt.Scanf("%d %d", w, h)

	return [][]int{}, nil
}
