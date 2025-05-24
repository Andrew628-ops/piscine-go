package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseRunes := []rune(base)
	baseLen := int64(len(baseRunes))

	var num int64 = int64(nbr)

	if num < 0 {
		z01.PrintRune('-')
		// Special handling for MinInt64
		if num == -9223372036854775808 {
			// Convert manually since abs overflows
			printBase(9223372036854775808, baseRunes, baseLen)
			return
		}
		num = -num
	}

	printBase(num, baseRunes, baseLen)
}

func printBase(num int64, base []rune, baseLen int64) {
	if num >= baseLen {
		printBase(num/baseLen, base, baseLen)
	}
	z01.PrintRune(base[num%baseLen])
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
