package piscine

func Rot14(s string) string {
	ls := []rune(s)
	for i := 0; i < len(ls); i++ {
		if ls[i] >= 'A' && ls[i] <= 'L' || ls[i] >= 'a' && ls[i] <= 'l' {
			ls[i] += 14
		} else if ls[i] >= 'M' && ls[i] <= 'Z' || ls[i] >= 'm' && ls[i] <= 'z' {
			ls[i] -= 12
		}
	}
	return string(ls)
}
