package main

import (
	"os"

	"github.com/01-edu/z01"
)

func validateOperator(test string) bool {
	op := []string{"+", "-", "*", "/", "%"}
	for _, res := range op {
		if res == test {
			return true
		}
	}
	return false
}

func atoi(s string) int {
	n := 0
	sign := 1
	for i, c := range s {
		if i == 0 && (c == '-' || c == '+') {
			if c == '-' {
				sign = -1
			}
			continue
		}
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c-'0')
	}
	return n * sign
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	if !validateOperator(args[1]) {
		return
	}

	premier := atoi(args[0])
	second := atoi(args[2])

	switch args[1] {
	case "+":
		z01.PrintRune(rune(premier + second))
		z01.PrintRune('\n')
	case "-":
		z01.PrintRune(rune(premier - second))
		z01.PrintRune('\n')
	case "*":
		z01.PrintRune(rune(premier * second))
		z01.PrintRune('\n')
	case "/":
		if second != 0 {
			z01.PrintRune(rune(premier / second))
			z01.PrintRune('\n')
		} else {
			for _, c := range "No division by 0\n" {
				z01.PrintRune(c)
			}
		}
	case "%":
		if second != 0 {
			z01.PrintRune(rune(premier % second))
			z01.PrintRune('\n')
		} else {
			for _, c := range "No modulo by 0\n" {
				z01.PrintRune(c)
			}
		}
	}
}
