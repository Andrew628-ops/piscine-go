package piscine 

import "github.com/01-edu/z01"

func PrintComb()  {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {

			PrintDigit(a / 10)
			PrintDigit(a % 10)
			z01.PrintRune('')

			PrintDigit(b / 10)
			PrintDigit(b % 10)

			if a != 98 || b != 99 {
				z01.PrintRune(',')
				z01.PrintRune('')
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintDigit(n int)  {
	digits := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	z01.PrintRune(digits[n])
}