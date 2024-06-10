package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run . <filename>")
		return
	}

	filename := os.Args[1]
	_, err := os.Open(filename)
	if err != nil {
		println("ERROR:", err)
		os.Exit(1)
	}
}
