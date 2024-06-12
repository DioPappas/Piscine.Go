package piscine

import "fmt"

// PrintNumber prints a single integer
func PrintNumber(n int) {
	fmt.Print(n)
}

// ForEach applies a given function to each element of an integer slice
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
