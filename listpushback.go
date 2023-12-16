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
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		// If the list is empty, set the new node as both the head and tail.
		l.Head = newNode
		l.Tail = newNode
	} else {
		// If the list is not empty, append the new node to the tail.
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
