package piscine

func Unmatch(a []int) int {
	result := 0

	// XOR all elements in the slice
	for _, num := range a {
		result ^= num
	}

	return result
}
