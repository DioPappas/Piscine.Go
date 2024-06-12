package piscine

// IsSorted checks if the given slice is sorted according to the provided comparison function.
// The comparison function f should return a negative value if a < b,
// zero if a == b, and a positive value if a > b.
func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	for i := 1; i < n; i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
