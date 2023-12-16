package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Perform ASCII sorting using Insertion Sort
	for i := 1; i < len(args); i++ {
		key := args[i]
		j := i - 1

		for j >= 0 && args[j] > key {
			args[j+1] = args[j]
			j = j - 1
		}
		args[j+1] = key
	}

	// Print the sorted arguments
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
