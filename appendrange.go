package piscine

func AppendRange(min, max int) []int {
	var dion []int
	for i := min; i < max; i++ {
		dion = append(dion, i)
	}
	return dion
}
