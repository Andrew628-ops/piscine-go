package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	var result []rune
	for nbr > 0 {
		remainder := nbr % baseLen
		result = append([]rune{rune(base[remainder])}, result...)
		nbr = nbr / baseLen
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
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
