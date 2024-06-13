package piscine

func StringToIntSlice(str string) []int {
	// Initialize an empty slice to store the ASCII values
	intSlice := []int{}

	// Iterate over each character in the string
	for _, char := range str {
		// Append the ASCII value of the character to the slice
		intSlice = append(intSlice, int(char))
	}

	// Return the slice
	return intSlice
}
