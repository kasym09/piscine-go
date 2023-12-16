package piscine

// CountIf returns the number of elements in a string slice for which the
// f function returns true.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, value := range tab {
		if f(value) {
			count++
		}
	}
	return count
}
