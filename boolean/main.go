package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(args []string) bool {
	if len(args)%2 == 0 {
		return true
	}
	return false
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	args := os.Args[1:]
	if isEven(args) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
