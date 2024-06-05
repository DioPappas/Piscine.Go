package piscine

func isAlphaNumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	ar := []rune(s)
	letter := true

	for i := 0; i < len(s); i++ {
		if isAlphaNumeric(ar[i]) && letter {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] -= 'a' - 'A'
			}
			letter = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] += 'a' - 'A'
		} else if !isAlphaNumeric(ar[i]) {
			letter = true
		}
	}

	return string(ar)
}
