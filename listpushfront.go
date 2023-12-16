package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		// If the list is empty, set the new node as both the head and tail.
		l.Head = newNode
		l.Tail = newNode
	} else {
		// If the list is not empty, insert the new node at the beginning.
		newNode.Next = l.Head
		l.Head = newNode
	}
}
