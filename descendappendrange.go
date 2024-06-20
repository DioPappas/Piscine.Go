package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{} // Return empty slice if max <= min
	}

	size := max - min           // Calculate the size of the slice
	result := make([]int, size) // Initialize a slice with 'size' capacity

	// Populate the slice with values from max down to min+1
	for i := 0; i < size; i++ {
		result[i] = max - i - 1
	}

	return result
}
