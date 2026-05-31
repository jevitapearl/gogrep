package search

import (
	"bufio"
	"gogrep/internal/config"
	"gogrep/internal/model"
	"os"
	"path/filepath"
	"regexp"
)

func SearchFile(fp, pattern string, opts config.Options, matchHandler func(model.Match)) error {
	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()

	filename := filepath.Base(fp)

	scanner := bufio.NewScanner(file)

	lineIndex := 0

	if opts.IgnoreCase {
		pattern = `(?i)` + pattern
	}

	if opts.WordMatch {
		pattern = `\b(?:` + pattern + `)\b`
	}

	if opts.LineMatch {
		pattern = `^(?:` + pattern + `)$`
	}

	searchPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	matchesEmptyLine := searchPattern.MatchString("")

	for scanner.Scan() {

		line := scanner.Text()
		lineIndex++

		if line == "" && !matchesEmptyLine {
			continue
		}

		col := searchPattern.FindStringIndex(line)

		matched := col != nil

		if opts.Invert {
			matched = !matched
		}

		column := 0
		if col != nil {
			column = col[0] + 1
		}

		if matched {
			results := model.Match{
				File:   filename,
				Line:   lineIndex,
				Column: column,
				Text:   line,
			}
			matchHandler(results)
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
