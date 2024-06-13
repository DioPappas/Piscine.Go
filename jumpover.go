package piscine

func JumpOver(str string) string {
	// Check if the string is empty or has fewer than 3 characters
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}

	// Initialize an empty string to store the result
	result := ""

	// Iterate through the input string with a step of 3
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// Append newline at the end as required
	result += "\n"

	return result
}
