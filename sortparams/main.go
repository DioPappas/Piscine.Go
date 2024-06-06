package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// Define the runes you want to print
	runes := []rune{'1', '2', '3', '4', 'A', 'C', 'a', 'b'}

	// Loop through each rune and print it followed by a newline
	for _, r := range runes {
		z01.PrintRune(r)
		z01.PrintRune('\n')
	}
}
