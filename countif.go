package piscine

// CountIf returns the number of elements of a string slice for which the f function returns true
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) {
			count++
		}
	}
	return count
}
