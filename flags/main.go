package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// If no arguments or --help flag provided, print usage instructions
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printUsage()
		return
	}

	var (
		insert string
		order  bool
		text   string
	)

	// Parse flags and arguments
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--insert", "-i":
			if i+1 < len(args) {
				insert = args[i+1]
				i++
			}
		case "--order", "-o":
			order = true
		default:
			text += args[i]
		}
	}

	// Concatenate the insertion string if provided and the insert flag is present
	if insert != "" && !order {
		text += insert
	}

	// Order the string if requested
	if order {
		text = orderString(text)
	}

	// Concatenate the insertion string if provided and the insert flag is not present
	if insert != "" && order {
		text = insert + text
	}

	// Print the resulting string
	for _, ch := range text {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  --insert (-i) <string>    \tInsert the given string into the argument string.")
	fmt.Println("  --order  (-o)             \tOrder the argument string in ASCII order.")
	fmt.Println("  --help   (-h)             \tPrint usage instructions.")
}

func orderString(s string) string {
	// Convert the string to a slice of runes
	runes := []rune(s)

	// Sort the runes
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}

	// Convert the sorted runes back to a string
	return string(runes)
}
