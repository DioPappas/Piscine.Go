package piscine

// Any returns true if at least one element in the string slice, when passed through the f function, returns true
func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}
