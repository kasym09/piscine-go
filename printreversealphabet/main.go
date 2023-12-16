package main

import (
	"github.com/01-edu/z01"
)

func main() {
	reverseAlphabet := "zyxwvutsrqponmlkjihgfedcba"

	for _, char := range reverseAlphabet {
		z01.PrintRune(char)
	}

	z01.PrintRune(10)
}
