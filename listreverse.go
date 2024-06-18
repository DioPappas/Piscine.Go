package piscine

func ListReverse1(l *List) {
	current := l.Head
	var prev *NodeL

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
