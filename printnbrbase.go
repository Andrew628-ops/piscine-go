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
	num := int64(nbr)

	if num < 0 {
		z01.PrintRune('-')
		// Instead of -num, use this trick:
		// Convert to unsigned to safely represent MinInt64
		printBase(uint64(-num), baseRunes, baseLen)
	} else {
		printBase(uint64(num), baseRunes, baseLen)
	}
}

func printBase(num uint64, base []rune, baseLen int64) {
	if num >= uint64(baseLen) {
		printBase(num/uint64(baseLen), base, baseLen)
	}
	z01.PrintRune(base[num%uint64(baseLen)])
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
