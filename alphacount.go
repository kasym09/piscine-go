package piscine

func AlphaCount(s string) int {
	counter := 0
	bytes := []byte(s)
	for _, i := range bytes {
		if i >= 'A' && i <= 'Z' || i >= 'a' && i <= 'z' {
			counter++
		}
	}
	return counter
}
