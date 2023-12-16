package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	massivRU := []rune(s)
	return massivRU[n-1]
}
