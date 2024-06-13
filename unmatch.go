package piscine

func Unmatch(a []int) int {
	count := make(map[int]int)
	for _, num := range a {
		count[num]++
	}
	for num, cnt := range count {
		if cnt%2 != 0 {
			return num
		}
	}

	return -1
}
