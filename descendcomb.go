package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	flagfirst := true
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '1'; j-- {
			for k := '9'; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if flagfirst == true {
						l = '8'
						flagfirst = false
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						if !(i == '0' && j == '1' && k == '0' && l == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						l--
					}
					if (i > k) || (i == k && j >= l) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						if !(i == '0' && j == '1' && k == '0' && l == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
