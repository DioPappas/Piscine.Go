package main

import (
	"os"
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

func atoi(s string) (int, bool) {
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
			return 0, false // Return error if not a valid number
		}
		n = n*10 + int(c-'0')
	}
	return n * sign, true
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	s := ""
	for n > 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	return sign + s
}

func writeToStdout(str string) {
	b := []byte(str)
	_, _ = os.Stdout.Write(b)
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	if !validateOperator(args[1]) {
		return
	}

	premier, ok1 := atoi(args[0])
	second, ok2 := atoi(args[2])

	if !ok1 || !ok2 {
		return // Exit if conversion fails
	}

	switch args[1] {
	case "+":
		writeToStdout(itoa(premier+second) + "\n")
	case "-":
		writeToStdout(itoa(premier-second) + "\n")
	case "*":
		writeToStdout(itoa(premier*second) + "\n")
	case "/":
		if second != 0 {
			writeToStdout(itoa(premier/second) + "\n")
		} else {
			writeToStdout("No division by 0\n")
		}
	case "%":
		if second != 0 {
			writeToStdout(itoa(premier%second) + "\n")
		} else {
			writeToStdout("No modulo by 0\n")
		}
	}
}
