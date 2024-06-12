package main

import (
	"os"
	"strconv"
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

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		// Do nothing
	} else {
		if validateOperator(args[1]) == false {
			println(0)
		} else {
			premier, _ := strconv.Atoi(args[0])
			second, _ := strconv.Atoi(args[2])

			if args[1] == "%" && second == 0 {
				println("No Modulo by 0")
			} else if args[1] == "/" && second == 0 {
				println("No division by 0")
			} else if args[1] == "+" {
				println(premier + second)
			} else if args[1] == "-" {
				println(premier - second)
			} else if args[1] == "*" {
				println(premier * second)
			} else if args[1] == "/" {
				println(premier / second)
			} else {
				println(premier % second)
			}
		}
	}
}
