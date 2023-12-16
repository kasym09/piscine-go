package piscine

// ForEach applies a function on each element of an int slice.
func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
