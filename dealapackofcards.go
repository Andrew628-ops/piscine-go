package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		playerHeader := fmt.Sprintf("Player %d: ", i+1)
		for _, r := range playerHeader {
			z01.PrintRune(r)
		}

		for j := 0; j < 3; j++ {
			card := fmt.Sprintf("%d", deck[i*3+j])
			for _, r := range card {
				z01.PrintRune(r)
			}
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
