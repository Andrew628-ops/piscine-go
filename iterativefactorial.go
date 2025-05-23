package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
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

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		if result > 0 && result > (1<<31-1)/i {
			return 0
		}
		result *= i
	}
	return result
}
