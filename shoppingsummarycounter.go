package piscine

// ShoppingSummaryCounter counts the occurrences of each item in the input string
func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	itemCount := make(map[string]int)

	// Initialize variables for processing the string
	currentItem := ""
	for _, char := range str {
		if char == ' ' {
			// If we encounter a space and currentItem is not empty, count the item
			if currentItem != "" {
				itemCount[currentItem]++
				currentItem = ""
			}
		} else {
			// Append the character to the current item
			currentItem += string(char)
		}
	}

	// Count the last item if there is any
	if currentItem != "" {
		itemCount[currentItem]++
	}

	return itemCount
}
