package flags

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func FlagHandler(flags []string) (map[string]int, error) {
	options := map[string]int{}

	for _, flag := range flags {
		key, err := validFlag(flag)

		if err != nil {
			log.Fatal(err)
		}

		if options[key] == 1 {
			log.Fatal(errors.New("repeated flags"))
		}

		if strings.HasPrefix(key, "delay-ms=") && strings.HasPrefix(key, "random=") {
			key, number, err := numericValue(key)

			if err != nil {
				log.Fatal(err)
			}

			options[key] = number
		} else {
			options[key] = 1
		}
	}

	return options, nil
}

func validFlag(s string) (string, error) {
	flags := []string{
		"--verbose",
		"--delay-ms=X",
		"--file=X",
		"--edges-portal",
		"--random=WxH",
		"--fullscreen",
		"--forprints",
		"--colored",
	}

	for _, r := range flags {
		if s == r {
			return r[2:], nil
		}
	}

	return "", errors.New("invalid flag")
}

func numericValue(key string) (string, int, error) {
	values := strings.Split(key, "=")

	key = values[0]
	number, err := strconv.Atoi(values[1])

	if err != nil {
		return "", 0, errors.New("unexpected string, while expecting int")
	}
	return key, number, nil
}
