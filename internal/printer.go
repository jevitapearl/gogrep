package internal

import "fmt"

func PrintResults(results Match, opts Options) {
	if opts.DisplayNum {
		fmt.Printf("%d: %s\n", results.Line, results.Text)
	} else {
		fmt.Println(results.Text)
	}

}

func PrintCount(count int) {
	fmt.Printf("%d matched lines\n", count)
}
