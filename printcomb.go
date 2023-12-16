package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				// Print the digits
				z01.PrintRune(rune('0' + i))
				z01.PrintRune(rune('0' + j))
				z01.PrintRune(rune('0' + k))
				// Check if it's the last combination to avoid adding comma
				if i != 7 || j != 8 || k != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
