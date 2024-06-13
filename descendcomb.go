package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		startJ := '9'
		if i == '9' {
			startJ = '8' // Start j from '8' when i is '9'
		}
		for j := startJ; j >= '0'; j-- {
			startK := '9'
			if i == '9' && j == '8' {
				startK = '8' // Start k from '8' when i is '9' and j is '8'
			}
			for k := startK; k >= '0'; k-- {
				startL := '9'
				if i == '9' && j == '8' && k == '8' {
					startL = '8' // Start l from '8' when i is '9', j is '8', and k is '8'
				}
				for l := startL; l >= '0'; l-- {
					if (i > k) || (i == k && j >= l) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')

						z01.PrintRune(k)
						z01.PrintRune(l)
						if !(i == '9' && j == '8' && k == '8' && l == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
