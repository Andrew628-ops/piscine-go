package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				// Convert digits to rune once, no cast at print time
				di := '0' + i
				dj := '0' + j
				dk := '0' + k

				z01.PrintRune(di)
				z01.PrintRune(dj)
				z01.PrintRune(dk)

				if !(i == 7 && j == 8 && k == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
