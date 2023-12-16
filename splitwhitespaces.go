package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	word := ""

	for _, char := range s {
		if isWhiteSpace(char) {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}

func isWhiteSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}
