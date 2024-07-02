package io

import (
	"game/assets"
	"log"
)

func MaintainInput(flagsMap map[string]string) [][]int {
	if flagsMap["random"] == "false" {
		temp, err := InputHandler()

		if err != nil {
			log.Fatal(err)
		}

		return temp
	}
	return assets.RandomGrid()
}

func InputHandler() ([][]int, error) {
	return [][]int{}, nil
}
