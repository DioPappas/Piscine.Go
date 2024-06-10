package piscine

import (
	"github.com/01-edu/z01"
)

// PrintNbr prints an int passed as a parameter
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	printPositive(n)
}

func printPositive(n int) {
	if n == 0 {
		return
	}
	printPositive(n / 10)
	z01.PrintRune(rune('0' + n%10))
}
