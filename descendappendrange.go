package piscine

func DescendAppendRange(max, min int) []int {
	// Check if max is less than or equal to min
	if max <= min {
		return []int{} // Return empty slice
	}

	// Calculate the number of elements in the range
	size := max - min

	// Create a slice with size elements
	result := make([]int, size)

	// Populate the slice in descending order
	for i := 0; i < size; i++ {
		result[i] = max - i - 1
	}

	return result
}
