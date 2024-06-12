// piscine/foreach.go
package piscine

import (
	"os"
	"strconv"
)

// ForEach applies a given function to each element of an integer slice
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// Directly prints a single integer without using fmt
func PrintInt(n int) {
	os.Stdout.WriteString(strconv.Itoa(n))
}
