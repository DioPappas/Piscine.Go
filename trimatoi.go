package piscine

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	found := false

	for _, char := range s {
		if char == '-' && !found {
			sign = -1
			found = true
		} else if isDigit(char) {
			found = true
			result = result*10 + int(char-'0')
		} else if found {
			break
		}
	}

	return result * sign
}
