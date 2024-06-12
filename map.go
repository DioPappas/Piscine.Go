package piscine

// Map applies a function of type func(int) bool on each element of an int slice and returns a slice of all the return values
func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}
