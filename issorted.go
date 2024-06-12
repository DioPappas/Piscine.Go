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

// Example comparison function: Sorts the integers in ascending order.
func isSortedByAscending(a, b int) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	}
	return 1
}

// Example usage:
// isSorted := IsSorted(isSortedByAscending, []int{3, 5, 7, 9, 11})
