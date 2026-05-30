package internal

import (
	"errors"
	"fmt"
	"os"
)

func GetArgs() (Options, string, string, error) {

	allArgs := os.Args
	var opts Options

	if len(allArgs) < 3 {
		return opts, "", "", errors.New("Not enough arguements")
	}

	if len(allArgs) == 3 {
		return opts, allArgs[1], allArgs[2], nil
	}

	flag := allArgs[1]
	switch flag {
	case "-i":
		opts.IgnoreCase = true

	case "-v":
		opts.Invert = true

	case "-c":
		opts.MatchCount = true

	default:
		fmt.Println("unknown flag")
		os.Exit(1)

	}

	if len(allArgs) == 4 {
		return opts, allArgs[2], allArgs[3], nil
	}

	return opts, "", "", nil

}
