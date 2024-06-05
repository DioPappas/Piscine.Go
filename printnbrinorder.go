package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return // do nothing for negative numbers
	}

	// Count occurrences of each digit in the input number
	digitCounts := [10]int{}
	for n > 0 {
		digit := n % 10
		digitCounts[digit]++
		n /= 10
	}

	// Print digits in ascending order
	for i := 0; i < 10; i++ {
		for j := 0; j < digitCounts[i]; j++ {
			z01.PrintRune(rune(i) + '0')
		}
	}

	// Print a newline character at the end
	z01.PrintRune('\n')
}
