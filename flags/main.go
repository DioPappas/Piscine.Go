package main

import (
	"fmt"
	"os"
	"sort"

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
			text += args[i] + " "
		}
	}

	// Insert the string if provided
	text += insert

	// Order the string if requested
	if order {
		text = orderString(text)
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
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
