package flags

import (
	"errors"
	"fmt"
	"game/io"
	"log"
	"os"
	"strconv"
	"strings"
)

func FlagHandler() map[string]interface{} {
	var err error
	var flagMap map[string]interface{}
	args := os.Args[1:]
	if len(args) != 0 {
		if args[0] == "--help" {
			if len(args) != 1 {
				log.Fatal("unexpected flag after --help")
			}
			io.Help()
			os.Exit(0)
		}
	}
	flagMap, err = flagsHandler(args)

	if err != nil {
		log.Fatal(err)
	}

	if flagMap["file"] != false && flagMap["random"] != false {
		log.Fatal("--file= and --random=WxH cannot be executed at the same time")
	}

	return flagMap
}

func flagsHandler(flags []string) (map[string]interface{}, error) {
	options := map[string]interface{}{
		"verbose":      false,
		"delay-ms":     2500,
		"file":         false,
		"edges-portal": false,
		"random":       false,
		"fullscreen":   false,
		"forprints":    false,
		"colored":      false,
	}

	for _, flag := range flags {
		key, err := validFlag(flag)
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(flag, "--delay-ms=") {
			newKey, number, _, err := numericValue(flag)
			if err != nil {
				return nil, err
			}
			options[newKey] = number
		} else if strings.HasPrefix(flag, "--random=") {
			newKey, h, w, err := numericValue(flag)
			if err != nil {
				return nil, err
			}
			options[newKey] = []int{h, w}
		} else if strings.HasPrefix(flag, "--file=") {
			filename := strings.SplitN(flag, "=", 2)
			if len(filename[1]) == 0 {
				return nil, errors.New("invalid value declaration in flag")
			}
			options["file"] = filename[1]
		} else {
			options[key] = "true"
		}
	}

	return options, nil
}

func validFlag(s string) (string, error) {
	flags := []string{
		"--verbose",
		"--edges-portal",
		"--fullscreen",
		"--forprints",
		"--colored",
		"--delay-ms=",
		"--random=",
		"--file=",
	}

	for _, flag := range flags {
		if s == flag {
			return flag[2:], nil
		} else if strings.HasPrefix(s, flag) {
			return flag[2 : len(flag)-1], nil
		}
	}

	return "", errors.New("invalid flag")
}

func numericValue(key string) (string, int, int, error) {
	values := strings.SplitN(key, "=", 2)

	if len(values) != 2 {
		return "", 0, 0, errors.New("invalid value declaration in flag")
	}

	key = values[0][2:] // Remove the leading "--"

	if key == "random" {
		dimensions := strings.SplitN(values[1], "x", 2)
		if len(dimensions) != 2 {
			return "", 0, 0, errors.New("invalid dimensions for random flag")
		}
		width, err := strconv.Atoi(dimensions[0])
		if err != nil {
			return "", 0, 0, fmt.Errorf("unexpected value for width, expecting int: %w", err)
		}
		height, err := strconv.Atoi(dimensions[1])
		if err != nil {
			return "", 0, 0, fmt.Errorf("unexpected value for height, expecting int: %w", err)
		}
		if height <= 0 {
			return "", 0, 0, errors.New("height cannot be smaller or equal to 0")
		}
		if width <= 0 {
			return "", 0, 0, errors.New("width cannot be smaller or equal to 0")
		}
		return key, width, height, nil
	}

	delay, err := strconv.Atoi(values[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("unexpected value, expecting int: %w", err)
	}
	if delay <= 0 {
		return "", 0, 0, errors.New("delay cannot be smaller or equal to 0")
	}
	return key, delay, 0, nil
}
