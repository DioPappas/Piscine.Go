package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	// os.Args contains the command-line arguments passed to the program
	args := os.Args[1:] // Skip the first argument which is the program name

	// Sort the arguments in alphabetical order
	sort.Strings(args)

	// Loop through each sorted argument
	for _, arg := range args {
		// Loop through each character in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
