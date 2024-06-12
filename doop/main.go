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

func isNumeric(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func add(a, b string) string {
	var result []byte
	carry := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 10
		sum %= 10
		result = append([]byte{byte(sum + '0')}, result...)
	}

	return string(result)
}

func subtract(a, b string) string {
	if b[0] == '-' {
		return add(a, b[1:])
	}
	var result []byte
	borrow := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 {
		diff := int(a[i] - '0' - byte(borrow))
		borrow = 0
		if j >= 0 {
			diff -= int(b[j] - '0')
			j--
		}
		if diff < 0 {
			diff += 10
			borrow = 1
		}
		result = append([]byte{byte(diff + '0')}, result...)
		i--
	}

	// Trim leading zeros
	for len(result) > 1 && result[0] == '0' {
		result = result[1:]
	}

	return string(result)
}

func multiply(a, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}

	var result []byte
	m := len(a)
	n := len(b)
	products := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			digitA := int(a[i] - '0')
			digitB := int(b[j] - '0')
			products[i+j+1] += digitA * digitB
		}
	}

	carry := 0
	for i := len(products) - 1; i >= 0; i-- {
		sum := products[i] + carry
		carry = sum / 10
		sum %= 10
		result = append([]byte{byte(sum + '0')}, result...)
	}

	if result[0] == '0' {
		result = result[1:]
	}

	return string(result)
}

func divide(a, b string) string {
	quotient := ""
	remainder := a

	for len(remainder) >= len(b) {
		num := b
		div := 0
		for isLargerOrEqual(remainder, num) {
			remainder = subtract(remainder, num)
			div++
		}
		quotient += string(div + '0')
	}

	if quotient == "" {
		return "0"
	}

	return quotient
}

func modulo(a, b string) string {
	remainder := a

	for len(remainder) >= len(b) {
		num := b
		for isLargerOrEqual(remainder, num) {
			remainder = subtract(remainder, num)
		}
	}

	if remainder == "" {
		return "0"
	}

	return remainder
}

func isLargerOrEqual(a, b string) bool {
	if len(a) > len(b) {
		return true
	} else if len(a) < len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}

	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	if !validateOperator(args[1]) || !isNumeric(args[0]) || !isNumeric(args[2]) {
		return
	}

	var result string

	switch args[1] {
	case "+":
		result = add(args[0], args[2])
	case "-":
		result = subtract(args[0], args[2])
	case "*":
		result = multiply(args[0], args[2])
	case "/":
		if args[2] == "0" {
			result = "No division by 0"
		} else {
			result = divide(args[0], args[2])
		}
	case "%":
		if args[2] == "0" {
			result = "No modulo by 0"
		} else {
			result = modulo(args[0], args[2])
		}
	}

	os.Stdout.Write([]byte(result + "\n"))
}
