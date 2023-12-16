package piscine

func FirstRune(s string) rune {
	for i, letter := range s {
		if i == 0 {
			return letter
		}
	}
	return 0
}
