package internal

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Match struct {
	File   string
	Line   int
	Column int
	Text   string
}

type Options struct {
	Invert     bool
	IgnoreCase bool
	MatchCount bool
	DisplayNum bool
}

func SearchFile(fp, pattern string, opts Options) error {
	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()

	filename := filepath.Base(fp)

	scanner := bufio.NewScanner(file)

	lineIndex := 0
	matchCount := 0
	for scanner.Scan() {

		line := scanner.Text()
		lineIndex++
		searchLine := line
		searchPattern := pattern

		if opts.IgnoreCase {
			searchLine = strings.ToLower(line)
			searchPattern = strings.ToLower(pattern)
		}

		col := strings.Index(searchLine, searchPattern)

		matched := col != -1
		if opts.Invert {
			matched = !matched
		}

		if matched {
			results := Match{
				File:   filename,
				Line:   lineIndex,
				Column: col + 1,
				Text:   line,
			}
			if !opts.MatchCount {
				PrintResults(results, opts)
			}
			matchCount++

		}

	}
	if opts.MatchCount {
		PrintCount(matchCount)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
