package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	// Start the loop from index 1 to exclude the program name
	for _, arg := range arguments[1:] {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n') // Add a newline after each argument
	}
}
