package main

import (
	"fmt"
	"gogrep/internal"
	"os"
	"path/filepath"
)

func main() {
	opts, pattern, filename, err := internal.GetArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fp := filepath.Join(cwd, filename)	

	if err := internal.SearchFile(fp, pattern, opts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
