package piscine

func ListSize(l *List) int {
	count := 0
	current := l.Head

	// Traverse the linked list and count the elements.
	for current != nil {
		count++
		current = current.Next
	}

	return count
}
