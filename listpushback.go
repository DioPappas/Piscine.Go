package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		// The list is empty, so the new node becomes both the head and the tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// The list is not empty, append the new node to the end and update the tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
