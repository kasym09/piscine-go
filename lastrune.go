package piscine

func LastRune(s string) rune {
	var last rune
	for _, char := range s {
		last = char
	}
	return last
}
