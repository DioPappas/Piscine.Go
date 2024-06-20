package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}

	size := max - min
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = max - i - 1
	}

	return result
}
