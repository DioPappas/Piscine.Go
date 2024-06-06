package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		num, valid := myAtoi(arg)
		if !valid || num < 1 || num > 26 {
			z01.PrintRune(' ')
		} else {
			if upper {
				z01.PrintRune(rune('A' + num - 1))
			} else {
				z01.PrintRune(rune('a' + num - 1))
			}
		}
	}
	z01.PrintRune('\n')
}

func myAtoi(s string) (int, bool) {
	num := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		num = num*10 + int(c-'0')
	}
	return num, true
}
