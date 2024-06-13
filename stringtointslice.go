package piscine

func StringToIntSlice(str string) []int {
	// Explicitly initialize an empty slice to store the ASCII values
	strrune := []rune(str)
	var intSlice []int

	// Iterate over each character in the string
	for _, char := range strrune {
		// Append the ASCII value of the character to the slice
		intSlice = append(intSlice, (int(char)))
	}

	// Return the slice
	return intSlice
}
