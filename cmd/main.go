package main

import (
	"fmt"
	"gogrep/internal/cli"
	"gogrep/internal/model"
	"gogrep/internal/output"
	"gogrep/internal/search"
	"os"
	"path/filepath"
)

func main() {
	opts, pattern, filename, err := cli.GetArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fp := filepath.Join(cwd, filename)

	matchCount := 0
	matchHandler := func(results model.Match) {
		if !opts.MatchCount {
			output.PrintResults(results, opts)
		}
		matchCount++
	}

	if err := search.SearchFile(fp, pattern, opts, matchHandler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.MatchCount {
		output.PrintCount(matchCount)
	}

}
