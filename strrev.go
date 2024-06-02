package piscine

func StrRev(s string) string {
	runes := []rune(s)
	start := 0
	end := len(runes) - 1

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
	return string(runes)
}
