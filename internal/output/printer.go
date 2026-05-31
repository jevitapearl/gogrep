package output

import (
	"fmt"
	"gogrep/internal/config"
	"gogrep/internal/model"
)

func PrintResults(results model.Match, opts config.Options) {
	if opts.DisplayNum {
		fmt.Printf("%d: %s\n", results.Line, results.Text)
	} else {
		fmt.Println(results.Text)
	}

}

func PrintCount(count int) {
	fmt.Printf("%d matched lines\n", count)
}
