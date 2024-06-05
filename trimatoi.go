package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	foundNum := false

	for _, let := range s {
		if let >= '0' && let <= '9' {
			if !foundNum {
				foundNum = true
			}
			num = num*10 + int(let-'0')
		} else if let == '-' && !foundNum {
			sign = -1
		}
	}

	return num * sign
}
