package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := rune(0); i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			// convert i and j to digits as runes without casting
			z01.PrintRune('0' + i/10)
			z01.PrintRune('0' + i%10)

			z01.PrintRune(' ')

			z01.PrintRune('0' + j/10)
			z01.PrintRune('0' + j%10)

			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
