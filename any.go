package piscine

// Any returns true if, when a string slice is passed through an f function,
// at least one element returns true.
func Any(f func(string) bool, a []string) bool {
	for _, value := range a {
		if f(value) {
			return true
		}
	}
	return false
}
