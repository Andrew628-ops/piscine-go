package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	first := true
	for a := 99; a >= 1; a-- {
		for b := a - 1; b >= 0; b-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			printTwoDigit(a)
			z01.PrintRune(' ')
			printTwoDigit(b)
			first = false
		}
	}
}

func printTwoDigit(n int) {
	z01.PrintRune(rune(n/10 + '0')) // tens digit
	z01.PrintRune(rune(n%10 + '0')) // units digit
}
