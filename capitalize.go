package piscine

func Capitalize(s string) string {
	w := []byte(s)
	upperNext := true
	for i, char := range w {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			if upperNext && char >= 'a' {
				w[i] = char - 'a' + 'A'
			} else if !upperNext && char >= 'A' && char <= 'Z' {
				w[i] = char + 'a' - 'A'
			}
			upperNext = false
		} else {
			upperNext = true
		}
	}
	return string(w)
}
