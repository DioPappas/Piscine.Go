package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	result += "\n"
	return result
}
