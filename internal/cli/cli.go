package cli

import (
	"errors"
	"flag"
	"gogrep/internal/config"
)

func GetArgs() (config.Options, string, string, error) {
	var opts config.Options

	flag.BoolVar(&opts.IgnoreCase, "i", false, "ignore case")
	flag.BoolVar(&opts.DisplayNum, "n", false, "display line number")
	flag.BoolVar(&opts.Invert, "v", false, "invert search")
	flag.BoolVar(&opts.MatchCount, "c", false, "match lines")
	flag.BoolVar(&opts.LineMatch, "x", false, "entire line match")
	flag.BoolVar(&opts.Regex, "E", false, "regex match")
	flag.BoolVar(&opts.WordMatch, "w", false, "entire word match")

	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		return opts, "", "", errors.New("usage: gogrep [flags] pattern filename")
	}
	pattern := args[0]
	filename := args[1]

	return opts, pattern, filename, nil
}
