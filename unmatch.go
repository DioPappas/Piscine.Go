package piscine

func Unmatch(a []int) int {
	// Create a map to count occurrences of each element
	count := make(map[int]int)

	// Count occurrences of each element in the slice
	for _, num := range a {
		count[num]++
	}

	// Find the element with an odd count (unmatched element)
	for num, cnt := range count {
		if cnt%2 != 0 {
			return num
		}
	}

	// If all elements have pairs, return -1
	return -1
}
