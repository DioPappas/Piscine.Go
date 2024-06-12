// piscine/foreach.go
package piscine

import "os"

// PrintInt prints a single integer without using fmt or additional packages
func PrintInt(n int) {
	if n == 0 {
		os.Stdout.WriteString("0")
		return
	}

	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append([]byte{byte(n%10 + '0')}, digits...)
		n /= 10
	}

	os.Stdout.Write(digits)
}

// ForEach applies a given function to each element of an integer slice
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
