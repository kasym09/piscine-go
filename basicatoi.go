package piscine

func BasicAtoi(s string) int {
	x := 0
	for _, digit := range s {
		y := int(digit - '0')
		x = x*10 + y
	}
	return x
}
