package main

import (
	"game/assets/flags"
	"game/io"
	"log"
	"os"
)

func main() {
	// Flags handler
	args := os.Args[1:]
	if len(args) != 0 {
		if args[0] == "--help" {
			if len(args) != 1 {
				log.Fatal("unexpected flag after --help")
			}
			io.Help()
			os.Exit(0)
		}
		_, err := flags.FlagHandler(args)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Input handler

	// Infinite for loop
}
