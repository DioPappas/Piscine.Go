package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type List1 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		// The list is empty, so the new node becomes both the head and the tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// The list is not empty, insert the new node at the beginning
		newNode.Next = l.Head
		l.Head = newNode
	}
}
