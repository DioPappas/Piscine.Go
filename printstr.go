package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	//	for _, key := range s {
	//		z01.PrintRune(key)
	//	}
	//	z01.PrintRune('\n')

	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}
