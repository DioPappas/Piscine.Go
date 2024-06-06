package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// os.Args contains the command-line arguments passed to the program
	args := os.Args[1:] // Skip the first argument which is the program name

	// Implement bubble sort to sort the arguments
	n := len(args)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Loop through each sorted argument
	for _, arg := range args {
		// Loop through each character in the argument
		for _, r := range arg {
			z01.PrintRune(r) // Print each character
		}
		z01.PrintRune('\n') // Print a newline after each argument
	}
}
