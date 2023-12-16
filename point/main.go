package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	PrintStrInt(points.x, points.y)
	z01.PrintRune(10)
}

func PrintStrInt(x, y int) {
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printDigit(x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printDigit(y)
}

func printDigit(num int) {
	z01.PrintRune(rune('0' + num/10))
	z01.PrintRune(rune('0' + num%10))
}
