package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	// We will use two pointers technique to compact the slice in-place
	// i: pointer to the current position we are filling in the compacted slice
	// j: pointer to iterate through the original slice
	i := 0
	for j := 0; j < len(slice); j++ {
		if slice[j] != "" {
			// Move non-zero element to the front
			slice[i] = slice[j]
			i++
		}
	}
	// The compacted slice now contains non-zero elements at the start
	// We need to resize the slice to remove the trailing zeros
	*ptr = slice[:i]
	return i
}
