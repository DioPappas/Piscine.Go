package piscine

import (
	"github.com/01-edu/z01"
)

// PrintNbr prints an int passed as a parameter
func PrintNbr(n int64) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		if n == -9223372036854775808 { // Special case for the minimum int64 value
			z01.PrintRune('-')
			z01.PrintRune('9')
			n = 223372036854775808 // Set n to the positive overflow value of int64
		} else {
			z01.PrintRune('-')
			n = -n
		}
	}
	printPositive(n)
}

func printPositive(n int64) {
	if n == 0 {
		return
	}
	printPositive(n / 10)
	z01.PrintRune(rune('0' + n%10))
}
