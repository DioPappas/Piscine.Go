package piscine

// ShoppingSummaryCounter counts the occurrences of each item in the input string
func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	itemCount := make(map[string]int)

	// Initialize variables for processing the string
	currentItem := ""
	lastChar := ' ' // Initialize lastChar as a space to handle the start of the string

	// Iterate through each character in the string
	for _, char := range str {
		if char == ' ' && lastChar != ' ' {
			// If we encounter a space and currentItem is not empty, count the item
			if currentItem != "" {
				itemCount[currentItem]++
				currentItem = ""
			}
		} else if char != ' ' {
			// Append the character to the current item
			currentItem += string(char)
		}

		// Update lastChar to the current character
		lastChar = char
	}

	// Count the last item if there is any
	if currentItem != "" {
		itemCount[currentItem]++
	}

	// Count consecutive spaces as empty strings
	if lastChar == ' ' {
		itemCount[""]++
	}

	// Return the map with the counts
	return itemCount
}
