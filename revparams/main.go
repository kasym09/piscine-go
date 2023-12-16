package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i >= 1; i-- {
		args := arguments[i]
		for _, char := range args {
			z01.PrintRune(char)
		}
		z01.PrintRune(10)
	}
}
