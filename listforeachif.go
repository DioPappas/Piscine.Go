package piscine

func IsPositive1(node interface{}) bool {
	switch v := node.(type) {
	case int:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	case byte:
		return v > 0
	case string, rune:
		return false
	}
	return false
}

func IsNegative1(node interface{}) bool {
	switch v := node.(type) {
	case int:
		return v < 0
	case float32:
		return v < 0
	case float64:
		return v < 0
	case byte:
		return false
	case string, rune:
		return false
	}
	return false
}

func IsNumeric1(node interface{}) bool {
	switch node.(type) {
	case int, float32, float64, byte:
		return true
	case string, rune:
		return false
	}
	return false
}

func ListForEachIf(list []interface{}, f func(interface{}) interface{}, comp func(interface{}) bool) []interface{} {
	for i, node := range list {
		if comp(node) {
			list[i] = f(node)
		}
	}
	return list
}

func stringToOne(node interface{}) interface{} {
	if _, ok := node.(string); ok {
		return 1
	}
	return node
}

func isString(node interface{}) bool {
	_, ok := node.(string)
	return ok
}
