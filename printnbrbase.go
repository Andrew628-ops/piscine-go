package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	baseLen := int64(len(base))
	var num int64 = int64(nbr)

	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	var result []rune
	for num > 0 {
		remainder := num % baseLen
		result = append([]rune{rune(base[remainder])}, result...)
		num = num / baseLen
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
