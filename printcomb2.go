package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			// Print the first two-digit number
			z01.PrintRune(rune('0' + i/10))
			z01.PrintRune(rune('0' + i%10))
			z01.PrintRune(' ')

			// Print the second two-digit number
			z01.PrintRune(rune('0' + j/10))
			z01.PrintRune(rune('0' + j%10))

			// Check if it's the last combination to avoid adding comma
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
