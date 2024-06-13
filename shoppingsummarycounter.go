package piscine

// ShoppingSummaryCounter counts the occurrences of each item in the input string
func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	itemCount := make(map[string]int)

	// Initialize variables for processing the string
	currentItem := ""
	lastWasSpace := true // Start with true to handle spaces at the beginning

	// Iterate through each character in the string
	for _, char := range str {
		if char == ' ' {
			if !lastWasSpace {
				// If currentItem is not empty, count it as an item
				if currentItem != "" {
					itemCount[currentItem]++
				}
				currentItem = ""
			}
			// Mark that the last character was a space
			lastWasSpace = true
		} else {
			// Append the character to currentItem
			currentItem += string(char)
			// Mark that the last character was not a space
			lastWasSpace = false
		}
	}

	// Count the last item if there is any
	if currentItem != "" {
		itemCount[currentItem]++
	}

	// Count consecutive spaces as empty string
	if lastWasSpace {
		itemCount[""]++
	}

	return itemCount
}
