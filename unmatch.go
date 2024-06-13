package piscine

func Unmatch(a []int) int {
	result := 0
	for _, num := range a {
		result ^= num
	}

	return result
}
