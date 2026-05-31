package cli

import (
	"errors"
	"gogrep/internal/config"
	"os"
)

func GetArgs() (config.Options, string, string, error) {

	allArgs := os.Args
	var opts config.Options

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

	case "-E":
		opts.Regex = true

	case "-w":
		opts.WordMatch = true

	case "-x":
		opts.LineMatch = true

	default:
		return opts, "", "", errors.New("unknown flag")

	}

	if len(allArgs) == 4 {
		return opts, allArgs[2], allArgs[3], nil
	}

	return opts, "", "", errors.New("Only 4 args are supported")

}
