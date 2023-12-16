package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, char := range arguments {
		for i, ch := range char {
			if i >= 2 {
				z01.PrintRune(ch)
			}
		}
	}
	z01.PrintRune('\n')
}
