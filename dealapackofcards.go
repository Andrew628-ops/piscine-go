package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		printString("Player ")
		z01.PrintRune('1' + rune(i))
		printString(": ")
		for j := 0; j < 3; j++ {
			printNumber(deck[i*3+j])
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var digits []rune
	for n > 0 {
		digits = append([]rune{'0' + rune(n%10)}, digits...)
		n /= 10
	}
	for _, r := range digits {
		z01.PrintRune(r)
	}
}
