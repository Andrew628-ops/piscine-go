package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	player := 1
	i := 0
	for player <= 4 {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune('0' + rune(player))
		z01.PrintRune(':')
		z01.PrintRune(' ')
		count := 0
		for count < 3 {
			num := deck[i]
			if num >= 10 {
				z01.PrintRune('1')
				z01.PrintRune('0' + rune(num-10))
			} else {
				z01.PrintRune('0' + rune(num))
			}
			if count < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			count++
			i++
		}
		z01.PrintRune('\n')
		player++
	}
}
