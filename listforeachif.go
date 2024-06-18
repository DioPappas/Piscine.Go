package piscine

type NodeL2 struct {
	Data interface{}
	Next *NodeL
}

type List2 struct {
	Head *NodeL
}

func IsAlNode(node *NodeL) bool {
	// Implement the logic for IsAlNode if required.
	// Assuming it returns true for some condition on the node.
	return true // Placeholder implementation
}

func IsPositive_node(node *NodeL) bool {
	switch v := node.Data.(type) {
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

func IsNegative_node(node *NodeL) bool {
	switch v := node.Data.(type) {
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

func IsNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return true
	case string, rune:
		return false
	}
	return false
}

func ListForEachIf(l *List, f func(*NodeL), comp func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if comp(it) {
			f(it)
		}
		it = it.Next
	}
}
