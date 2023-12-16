package piscine

func TrimAtoi(s string) int {
	number := 0
	negative := false
	for _, char := range s {
		if char == '-' && !negative && number == 0 {
			negative = true
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		}
	}
	if negative {
		return -number
	}
	return number
}
