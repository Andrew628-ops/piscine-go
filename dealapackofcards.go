package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	player := 1
	index := 0
	for player <= 4 {
		printChar('P')
		printChar('l')
		printChar('a')
		printChar('y')
		printChar('e')
		printChar('r')
		printChar(' ')
		printChar(asciiDigit(player))
		printChar(':')
		printChar(' ')
		count := 0
		for count < 3 {
			n := deck[index]
			if n >= 10 {
				printChar('1')
				printChar(asciiDigit(n - 10))
			} else {
				printChar(asciiDigit(n))
			}
			if count < 2 {
				printChar(',')
				printChar(' ')
			}
			count++
			index++
		}
		z01.PrintRune('\n')
		player++
	}
}

func printChar(c byte) {
	z01.PrintRune(rune(c))
}

func asciiDigit(n int) byte {
	return byte('0' + n)
}
