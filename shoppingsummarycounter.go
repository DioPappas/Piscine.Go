package piscine

import (
	"strings"
)

// ShoppingSummaryCounter counts the occurrences of each item in the input string
func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	itemCount := make(map[string]int)

	// Split the input string into words based on spaces
	items := strings.Fields(str)

	// Iterate through the list of items and count each one
	for _, item := range items {
		itemCount[item]++
	}

	// Count consecutive spaces as empty strings
	consecutiveSpaces := strings.Count(str, " ")
	itemCount[""] += consecutiveSpaces

	// Return the map with the counts
	return itemCount
}
