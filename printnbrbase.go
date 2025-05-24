package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseRunes := []rune(base)
	baseLen := int64(len(baseRunes))

	var num int64 = int64(nbr)
	if num == 0 {
		z01.PrintRune(baseRunes[0])
		return
	}

	if num < 0 {
		z01.PrintRune('-')
		// Handle the special case of math.MinInt64
		// Convert to unsigned to prevent overflow
		printBase(-num, baseRunes, baseLen)
		return
	}

	printBase(num, baseRunes, baseLen)
}

func printBase(n int64, baseRunes []rune, baseLen int64) {
	if n >= baseLen {
		printBase(n/baseLen, baseRunes, baseLen)
	}
	z01.PrintRune(baseRunes[n%baseLen])
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
