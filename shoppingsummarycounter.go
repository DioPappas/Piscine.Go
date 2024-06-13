package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	itemCount := make(map[string]int)
	currentItem := ""
	for _, char := range str {
		if char == ' ' {
			if currentItem != "" {
				itemCount[currentItem]++
				currentItem = ""
			}
		} else {
			currentItem += string(char)
		}
	}

	if currentItem != "" {
		itemCount[currentItem]++
	}

	return itemCount
}
