package piscine

func Index(s string, toFind string) int {
	lenS := len(s)
	lenToFind := len(toFind)

	if lenToFind == 0 {
		return 0
	}

	for i := 0; i <= lenS-lenToFind; i++ {
		j := 0
		for j < lenToFind && s[i+j] == toFind[j] {
			j++
		}
		if j == lenToFind {
			return i
		}
	}

	return -1
}
