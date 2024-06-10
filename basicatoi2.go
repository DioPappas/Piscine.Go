package piscine

// BasicAtoi2 simulates the behavior of the Atoi function in Go
func BasicAtoi2(s string) int {
	var result int
	for _, char := range s {
		if char < '0' || char > '9' {
			// If the character is not a digit, return 0
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result
}
