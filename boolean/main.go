package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	numArgs := len(os.Args) - 1

	if numArgs%2 == 0 {
		printString("I have an even number of arguments")
	} else {
		printString("I have an odd number of arguments")
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
