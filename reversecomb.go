package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9' ; i >= '0'; i-- {
		for j := '9'; j >= '1' ; j-- {
			for k := '9'; k >= '0' ; k-- {
				for l := '8'; l >= '0' ; l-- {
					if (l < j)|| ( l == j && k < i )
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune()
					z01.PrintRune(k)
					z01.PrintRune(l)
					if !(i ='0' && j = '1'&& k = '0'&& l = '0' ) {
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
