package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	var totalLen int
	for _, arg := range args {
		totalLen += len(arg) + 1 // Add 1 for the newline character
	}

	result := make([]byte, totalLen-1) // Subtract 1 to exclude the newline after the last argument
	index := copy(result, args[0])

	for _, arg := range args[1:] {
		result[index] = '\n'
		index += copy(result[index+1:], arg) + 1
	}

	return string(result)
}
