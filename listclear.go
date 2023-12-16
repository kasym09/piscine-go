package piscine

func ListClear(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	// Iterate through the list and free each node
	current := l.Head
	for current != nil {
		nextNode := current.Next
		current.Next = nil // Clear the Next pointer to avoid accidental cyclic references
		current = nextNode
	}

	// Set the list's pointers to nil
	l.Head = nil
	l.Tail = nil
}
