package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// If no arguments are provided, read from stdin
		cat(os.Stdin)
		return
	}

	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Printf("ERROR: open %s: %v\n", arg, err)
			continue
		}
		cat(file)
		file.Close()
	}
}

func cat(reader io.Reader) {
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Print(string(content))
}
