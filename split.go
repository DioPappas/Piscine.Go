package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	count := 0

	// Iterate through the string s
	for i := 0; i < len(s); i++ {
		// Check if the separator occurs at position i
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			// Add the substring from start to i (excluding the separator) to the result slice
			if start != i { // Avoid adding empty strings
				// Ensure the result slice has enough capacity
				if result == nil {
					result = make([]string, 0, 1)
				} else if count >= cap(result) {
					newSlice := make([]string, len(result), 2*len(result)+1)
					copy(newSlice, result)
					result = newSlice
				}
				result = result[:count+1] // Expand the result slice by 1
				result[count] = s[start:i]
				count++
			}
			start = i + len(sep) // Move start to the next character after the separator
			i = start - 1        // Move i to the next character after the separator
		}
	}

	// Add the remaining substring to the result slice
	if start < len(s) {
		if result == nil {
			result = make([]string, 0, 1)
		} else if count >= cap(result) {
			newSlice := make([]string, len(result), 2*len(result)+1)
			copy(newSlice, result)
			result = newSlice
		}
		result = result[:count+1] // Expand the result slice by 1
		result[count] = s[start:]
	}

	return result
}
