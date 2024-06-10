package piscine

// Atoi simulates the behavior of the Atoi function in Go
func Atoi(s string) int {
	var result int
	var sign int = 1

	for i, char := range s {
		if i == 0 && (char == '+' || char == '-') {
			// If the first character is a sign, determine the sign
			if char == '-' {
				sign = -1
			}
			continue // Skip the sign character
		}
		if char < '0' || char > '9' {
			// If any character after the sign is not a digit, return 0
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result * sign
}
